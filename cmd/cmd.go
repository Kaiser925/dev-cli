package cmd

import "github.com/spf13/cobra"

// NewDefaultDevCtlCommand creates the `kubectl` command.
func NewDefaultDevCtlCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "devctl [command]",
		Long: "devctl build development environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}
