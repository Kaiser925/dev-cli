package cmd

import (
	"github.com/Kaiser925/devctl/resourses"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [resource kind]",
	Short: "Delete local resource",
	Long:  "Delete local resource, such as local mongo replica set.",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := parseConfig(resourceCfg, args)
		if err != nil {
			return err
		}
		return resourses.NewResourceOperator().DeleteResource(config)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&resourceCfg, "filename", "f", "",
		"that contains the configuration to createCmd")
}
