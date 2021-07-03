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
	//sequences.StaffGoddessChair(),
	sequences.Warriors(),
	//sequences.LowBackHipsGlutes().Action().Do()
	//sequences.WarriorSequence().Action().Do()
	//sequences.UpperBackShoulders().Action().Do()
}

func timer(minutes time.Duration, message string) {
	model.SpeakAndPause(
		fmt.Sprintf("starting timer for: %v minutes", int64(minutes)), 0).Action().Do()
	go func() {
		<-time.After(time.Minute * minutes)
		model.SpeakAndPause(
			fmt.Sprintf("time elapsed: %v minutes. %s", int64(minutes), message), 0).Action().Do()
	}()
}

func main() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	action.StartPauser(tty)

	timer(20, "wow you are doing really good")

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
