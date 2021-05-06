package action

import (
	"context"
	"fmt"
	"github.com/mattn/go-tty"
	"go.uber.org/atomic"
	"log"
	"os/exec"
	"sync"
	"time"
)

var Pauser actionPauser

type actionPauser struct {
	lock   sync.RWMutex
	pause  func()
	resume func()
	paused *atomic.Bool
	tty    *tty.TTY
}

func StartPauser(tty *tty.TTY) {
	Pauser = actionPauser{
		tty: tty,
		paused: atomic.NewBool(false),
	}
	go Pauser.run()
}

func (p *actionPauser) onPause(pause, resume func()) {
	p.lock.Lock()
	p.pause = pause
	p.resume = resume
	p.lock.Unlock()
}

func (p *actionPauser) waitUntilUnpaused() {
	for p.paused.Load() {
		time.Sleep(time.Millisecond * 50)
	}
}

func (p *actionPauser) run() {
	var paused bool
	for {
		r, err := p.tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		p.lock.RLock()
		pause := p.pause
		resume := p.resume
		p.lock.RUnlock()
		if pause == nil {
			SpeakAction{S: "no action, ignoring pause"}.Do()
			continue
		}
		paused = !paused
		fmt.Println("Paused => "+string(r), paused)
		p.paused.Store(paused)
		if paused {
			pause()
		} else {
			resume()
		}
	}
}

type Actions []Action

func (as Actions) Do() {
	for _, a := range as {
		a.Do()
	}
}

type Action interface {
	Do()
}

type PauseableAction interface {
	Action
	Pause()
	Resume()
}

type NullAction struct{}

func (a NullAction) Do() {}

func (a NullAction) Pause() {}

func (a NullAction) Resume() {}

type SpeakAction struct {
	S         string
	Cmd       *exec.Cmd
	pauseLock sync.Mutex
}

func (a SpeakAction) Do() {
	log.Printf("%v", a.S)
	cmd := exec.Command("espeak", a.S)
	a.Cmd = cmd
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	Pauser.onPause(
		a.Pause,
		a.Resume,
	)

	if err := a.Cmd.Wait(); err != nil {
		//log.Printf("is this an err?: %v", err)
	}

	Pauser.waitUntilUnpaused()
}

func (a SpeakAction) Pause() {
	for a.Cmd == nil || a.Cmd.Process == nil {
		time.Sleep(time.Millisecond * 50)
	}
	if err := a.Cmd.Process.Kill(); err != nil {
		log.Fatal(err)
	}
}

func (a SpeakAction) Resume() {
	// restart
	a.Do()
}

type SleepAction struct {
	D         time.Duration
	done      chan struct{}
	startTime time.Time
	remaining time.Duration
}

func (a SleepAction) Do() {
	log.Printf(">> sleep %v", a.D)
	if a.D < 0 {
		return
	}
	log.Printf("%v", a.D)
	tick := time.NewTicker(a.D)
	a.startTime = time.Now()
	a.done = make(chan struct{})
	Pauser.onPause(
		a.Pause,
		a.Resume,
	)
	// wait
	select {
	case <-a.done:
	case <-tick.C:
	}
	<-tick.C
	Pauser.waitUntilUnpaused()
}

func (a SleepAction) Pause() {
	a.remaining = a.D - (time.Now().Sub(a.startTime))
	close(a.done)
}

func (a SleepAction) Resume() {
	// restart
	log.Printf("continuing sleep with %s", a.remaining)
	SleepAction{
		D: a.remaining,
	}.Do()
}

type LoopingAction struct {
	Ctx    context.Context
	Action Action
}

func (a LoopingAction) Do() {
	for {
		select {
		case <-a.Ctx.Done():
			return
		default:
			a.Action.Do()
		}
	}
}
