package cmd

import (
	"errors"
	"fintrack/cmd/cli"
	"fintrack/pkg/shutdown"
	"log"
	"os"
)

func Execute() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	cleanup, err := handleRun()
	defer cleanup()

	if err != nil {
		log.Println(err)
		exitCode = exitCodeFailure
		return
	}

	shutdown.Gracefully()
}

// handleRun is a function that handles the execution of a command.
//
// It takes no parameters and returns a function and an error.
func handleRun() (func(), error) {
	var command string

	if len(os.Args) < 1 {
		return nil, errors.New(errNoCommand)
	}

	command = os.Args[1]

	switch command {
	case "cli":
		log.SetPrefix(logPrefix)
		log.SetFlags(logFlags)
		return cli.Run()
	default:
		return nil, errors.New(errUnknownCommand)
	}
}
