package cmd

import (
	"errors"
	"github.com/Kaiser925/devctl/common"
	"github.com/spf13/cobra"
)

func configValidator(filename string) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 && len(filename) == 0 {
			return errors.New("requires resources kind")
		}
		return nil
	}
}

func configFromArgs(args []string) (*common.ResourceConfig, error) {
	config := common.NewResourceConfig()
	switch len(args) {
	case 1:
		config.Kind = args[0]
	case 4:
		config.Kind = args[0]
		config.DatabaseName = args[1]
		config.User = args[2]
		config.Password = args[3]
	default:
		return config, errors.New("parameters is not right")
	}
	return config, nil
}
