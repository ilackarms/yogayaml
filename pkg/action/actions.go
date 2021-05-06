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
	paused *atomic.Bool
	tty    *tty.TTY
}

func StartPauser(tty *tty.TTY) {
	Pauser = actionPauser{
		tty:    tty,
		paused: atomic.NewBool(false),
	}
	go Pauser.run()
}

func (p *actionPauser) onPause(pause func()) {
	p.lock.Lock()
	p.pause = pause
	p.lock.Unlock()
}

func (p *actionPauser) waitUntilUnpaused(resume func()) {
	var wasPaused bool
	for p.paused.Load() {
		time.Sleep(time.Millisecond * 50)
		wasPaused = true
	}
	if wasPaused {
		resume()
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
		p.lock.RUnlock()
		if pause == nil {
			SpeakAction{S: "no action, ignoring pause"}.Do()
			continue
		}
		paused = !paused
		fmt.Println("Paused => ", string(r), paused)
		p.paused.Store(paused)
		if paused {
			pause()
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

type NullAction struct{}

func (a NullAction) Do() {}

func (a NullAction) Pause() {}

func (a NullAction) Resume() {}

type SpeakAction struct {
	S string
}

func (a SpeakAction) Do() {
	log.Printf("%v", a.S)
	cmd := exec.Command("espeak", a.S)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	Pauser.onPause(
		func() {
			for cmd.Process == nil {
				time.Sleep(time.Millisecond * 50)
			}
			if err := cmd.Process.Kill(); err != nil {
				log.Fatal(err)
			}
		},
	)

	if err := cmd.Wait(); err != nil {
		//log.Printf("is this an err?: %v", err)
	}

	Pauser.waitUntilUnpaused(func() {
		// restart
		SpeakAction{S: "resumed: " + a.S}.Do()
	})
}

type SleepAction struct {
	D time.Duration
}

func (a SleepAction) Do() {
	log.Printf(">> sleep %v", a.D)
	if a.D <= 0 {
		return
	}
	started := time.Now()
	tick := time.NewTicker(a.D)
	paused := make(chan struct{})
	Pauser.onPause(
		func() {
			paused <- struct{}{}
		},
	)
	for {
		select {
		case <-paused:
			elapsed := time.Since(started)
			remaining := a.D - elapsed
			if remaining <= 0 {
				// done
				return
			}
			log.Printf("sleep: paused (%v elapsed, %v remaining)", elapsed, remaining)
			// wait for unpause
			Pauser.waitUntilUnpaused(func() {
				log.Printf("sleep: unpaused")
				tick = time.NewTicker(remaining)
			})
		case <-tick.C:
			log.Printf("sleep: done")
			// done
			return
		}
	}
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
