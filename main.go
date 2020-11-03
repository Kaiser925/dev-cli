package main

import (
	"github.com/Kaiser925/devctl/cmd"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:  "devctl [command]",
	Long: "devopsctls build env for local test",
}

func main() {
	rootCmd.AddCommand(cmd.CreateCmd)
	rootCmd.AddCommand(cmd.RemoveCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
