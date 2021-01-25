package main

import (
	"os"

	"github.com/ess/fbz/cmd/fbz/commands"
)

func main() {
	if commands.Execute() != nil {
		os.Exit(1)
	}
}
