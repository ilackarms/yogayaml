package model

import (
	"bytes"
	"context"
	"github.com/ilackarms/yogayaml/pkg/action"
	"html/template"
	"log"
	"time"
)

// nodes can be collapsed into their actions in the graph
type Node interface {
	// generate the action to run
	Action() action.Action
	// skip this step
	Skip() bool
}

// a node that does nothing
type NoOpNode struct{}

func (n NoOpNode) Action() action.Action {
	return action.NullAction{}
}

func (n NoOpNode) Skip() bool {
	return false
}

// used to restart a sequenec from a specific node
type SkipPrevious struct{}

func (n SkipPrevious) Action() action.Action {
	return action.NullAction{}
}

func (n SkipPrevious) Skip() bool {
	return false
}

type Sequence []Node

// Flatten to an action
func (s Sequence) Action() action.Action {
	var actions action.Actions

	// preprocess
	var startIndex int
	for i, sequence := range s {
		if _, isSkip := sequence.(SkipPrevious); isSkip {
			startIndex = i
		}
	}
	for i, sequence := range s {
		if i < startIndex {
			continue
		}
		if !sequence.Skip() {
			actions = append(actions, sequence.Action())
		}
	}
	return actions
}

func (n Sequence) Skip() bool {
	skip := true
	for _, sequence := range n {
		skip = skip && sequence.Skip()
	}
	return skip
}

type PoseNode struct {
	Intro SpeechNode `json:"intro"`
	Pause PauseNode  `json:"pause"`
	Outro SpeechNode `json:"outro"`
	Next  Node       `json:"next"`
}

func (s PoseNode) Action() action.Action {
	var actions action.Actions
	if !s.Intro.Skip() {
		actions = append(actions, s.Intro.Action())
	}
	if !s.Pause.Skip() {
		actions = append(actions, s.Pause.Action())
	}
	if !s.Outro.Skip() {
		actions = append(actions, s.Outro.Action())
	}
	if s.Next != nil && !s.Next.Skip() {
		actions = append(actions, s.Next.Action())
	}
	return actions
}

func (n PoseNode) Skip() bool {
	return n.Intro.Skip() && n.Pause.Skip() && n.Outro.Skip()
}

type SpeechNode struct {
	Template   string                 `json:"text"`
	Parameters map[string]interface{} `json:"parameters"`
	Next       Node                   `json:"next"`
}

func (n SpeechNode) Action() action.Action {
	buf := &bytes.Buffer{}
	err := template.Must(template.New("speech").Parse(n.Template)).Execute(buf, n.Parameters)
	if err != nil {
		log.Fatal(err)
	}
	actions := action.Actions{action.SpeakAction{S: buf.String()}}
	if n.Next != nil && !n.Next.Skip() {
		actions = append(actions, n.Next.Action())
	}
	return actions
}

func (n SpeechNode) Skip() bool {
	return n.Template == ""
}

type PauseNode struct {
	LoopingText         SpeechNode    `json:"loopingText"`
	LoopingTextInterval time.Duration `json:"loopingTextInterval"`
	Duration            time.Duration `json:"duration"`
	Next                Node          `json:"next"`
}

func (n PauseNode) Action() action.Action {

	var actions action.Actions

	if n.LoopingText.Template == "" {
		actions = append(actions, action.SleepAction{D: n.Duration})
	} else {

		loopTextOnce := action.Actions{
			n.LoopingText.Action(),
			action.SleepAction{D: n.LoopingTextInterval},
		}

		ctx, cancel := context.WithTimeout(context.Background(), n.Duration)
		defer cancel()

		actions = append(actions, action.LoopingAction{
			Ctx:    ctx,
			Action: loopTextOnce,
		})
	}
	if n.Next != nil && !n.Next.Skip() {
		actions = append(actions, n.Next.Action())
	}
	return actions
}

func (n PauseNode) Skip() bool {
	return n.Duration == 0
}

type LoopNode struct {
	Node  Node `json:"node"`
	Times int  `json:"times"`
	Next  Node `json:"next"`
}

func (n LoopNode) Action() action.Action {

	var actions action.Actions

	for i := 0; i < n.Times; i++ {
		actions = append(actions, n.Node.Action())
	}

	if n.Next != nil && !n.Next.Skip() {
		actions = append(actions, n.Next.Action())
	}
	return actions
}

func (n LoopNode) Skip() bool {
	return n.Node.Skip()
}
