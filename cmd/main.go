package main

import (
	"fmt"
	"github.com/ilackarms/yogayaml/pkg/action"
	"github.com/ilackarms/yogayaml/pkg/model"
	"github.com/ilackarms/yogayaml/pkg/sequences"
	"github.com/mattn/go-tty"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Sequences = model.Sequence{
	sequences.FunSimple(),
	//sequences.LowBackHipsGlutes().Action().Do()
	//sequences.WarriorSequence().Action().Do()
	//sequences.UpperBackShoulders().Action().Do()
}

func main() {

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	action.StartPauser(tty)

	started := time.Now()
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		Sequences.Action().Do()
		close(done)
	}()

	<-done
	action.SpeakAction{S: fmt.Sprintf("finished: %v", time.Now().Sub(started))}.Do()
}
