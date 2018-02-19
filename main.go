package main

import (
	"log"

	"github.com/uniplaces/logfairy/cmd"
)

func main() {
	rootCmd := cmd.GetCommand(cmd.Setup())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
