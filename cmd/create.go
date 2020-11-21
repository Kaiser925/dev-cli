package cmd

import (
	"github.com/Kaiser925/devctl/resourses"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [resource kind]",
	Short: "create local resource",
	Long:  "create local resource, such as local mongo replica set.",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := parseConfig(resourceCfg, args)
		if err != nil {
			return err
		}
		return resourses.NewResourceOperator().CreateResource(config)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&resourceCfg, "filename", "f", "",
		"that contains the configuration to createCmd")
}
