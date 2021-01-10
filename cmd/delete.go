/*
 * Developed by Kaiser925 on 2020/12/17.
 * Lasted modified 2020/12/17.
 * Copyright (c) 2020.  All rights reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/Kaiser925/dev-cli/pkg/resourses"
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
