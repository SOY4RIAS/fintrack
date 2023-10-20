package cmd

import "log"

const (

	// exitCodeFailure is the exit code for a failed execution of the program.
	exitCodeFailure = 1

	// logPrefix is the prefix for log messages.
	logPrefix = string("FTRK CLI: ")

	// logFlags is the flags for log messages.
	logFlags = int(log.Lmsgprefix)

	// errUnknownCommand is the error message for an unknown command.
	errUnknownCommand = string("unknown command")

	// errNoCommand is the error message for no command provided.
	errNoCommand = string("no command provided")
)
