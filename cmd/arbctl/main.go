package main

import (
	"os"

	"github.com/alt-research/arbitrum-orbit-sdk-go/cmd/arbctl/commands"
	"github.com/urfave/cli"
)

func main() {
	err := commands.Run()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error() + "\n")
		exitCode := 1
		if exitCoder, ok := err.(cli.ExitCoder); ok {
			exitCode = exitCoder.ExitCode()
		}
		os.Exit(exitCode)
	}
}
