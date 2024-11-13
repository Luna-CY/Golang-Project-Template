package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/cmd/main/command"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	time.FixedZone("UTC+8", 8*60*60) // Set the timezone to UTC+8

	var wg sync.WaitGroup
	var ch = make(chan os.Signal, 1)

	// Create a context with cancellation
	var ctx, cancel = context.WithCancel(context.Background())

	// Register signal handlers
	signal.Notify(ch, os.Interrupt, syscall.SIGUSR1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := command.NewMainCommand().ExecuteContext(ctx); nil != err {
			if !errors.Is(err, context.Canceled) {
				_, _ = fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
			}
		}
	}()

	<-ch

	// Handle interrupt signal by canceling the context and exiting the program
	cancel()

	// Wait for the main command to finish
	wg.Wait()
}
