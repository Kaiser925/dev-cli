package cmd

import (
	"github.com/Kaiser925/devctl/common"
	"github.com/Kaiser925/devctl/resourses"
	"github.com/spf13/cobra"
)

var Create = &cobra.Command{
	Use:   "create <resource kind>",
	Short: "create local resource",
	Long:  "create local resource, such as local mongo replica set.",
	Args:  configValidators,
	RunE: func(cmd *cobra.Command, args []string) error {
		var config *common.ResourceConfig
		var err error
		if len(configfile) > 0 {
			config, err = common.ReadConfigFromFile(configfile)
		} else {
			config, err = configFromArgs(args)
		}

		if err != nil {
			return err
		}

		return resourses.NewResourceOperator().CreateResource(config)
	},
}

func init() {
	Create.Flags().StringVarP(&configfile, "filename", "f", "",
		"that contains the configuration to create")
}
