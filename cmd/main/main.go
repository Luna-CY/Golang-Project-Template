package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/cmd/main/command"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	time.FixedZone("UTC+8", 8*60*60) // Set the timezone to UTC+8

	var ctx, cancel = context.WithCancel(context.Background())

	var ch = make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGUSR1)
	go func() {
		<-ch
		cancel()
	}()

	if err := command.NewMainCommand().ExecuteContext(ctx); nil != err {
		if !errors.Is(err, context.Canceled) {
			_, _ = fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		}
	}
}
