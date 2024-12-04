package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/cmd/main/command"
	"os"
	"time"
)

func main() {
	time.FixedZone("UTC+8", 8*60*60) // Set the timezone to UTC+8

	if err := command.NewMainCommand().ExecuteContext(context.Background()); nil != err {
		if !errors.Is(err, context.Canceled) {
			_, _ = fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		}
	}
}
