package cmd

import (
	"errors"
	"github.com/Kaiser925/devctl/common"
	"github.com/spf13/cobra"
	"log"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:  "devctl [command]",
	Long: "devctl build env for local test",
}

var resourceCfg string

func parseConfig(resourceCfg string, args []string) (*common.ResourceConfig, error) {
	if len(resourceCfg) == 0 && len(args) == 0 {
		return nil, errors.New("missing config")
	}

	if len(resourceCfg) > 0 {
		return common.ReadConfigFromFile(resourceCfg)
	} else {
		return parseConfigFromArgs(args)
	}
}

func parseConfigFromArgs(args []string) (*common.ResourceConfig, error) {
	config := common.NewResourceConfig()
	switch len(args) {
	case 1:
		config.Kind = args[0]
	default:
		return config, errors.New("parameters is not right")
	}
	return config, nil
}
