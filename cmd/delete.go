package cmd

import (
	"github.com/Kaiser925/devctl/common"
	"github.com/Kaiser925/devctl/resourses"
	"github.com/spf13/cobra"
)

var Delete = &cobra.Command{
	Use:   "delete <resource>",
	Short: "Delete local resource",
	Long:  "Delete local resource, such as local mongo replica set.",
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

		return resourses.NewResourceOperator().DeleteResource(config)
	},
}

func init() {
	Delete.Flags().StringVarP(&configfile, "filename", "f", "",
		"that contains the configuration to create")
}
