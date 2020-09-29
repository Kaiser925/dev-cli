package main

import (
	"github.com/Kaiser925/devctl/cmd"
	"log"
)

func main() {
	command := cmd.NewDefaultDevCtlCommand()

	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
