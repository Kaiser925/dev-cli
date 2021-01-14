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
	"errors"
	"github.com/Kaiser925/dev-cli/pkg/common"
	"github.com/spf13/cobra"
	"log"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:  "dev-cli [command]",
	Long: "dev-cli build env for local test",
}

var resourceCfg string

func parseConfig(resourceCfg string, args []string) (*common.ResourceConfig, error) {
	if len(resourceCfg) == 0 && len(args) == 0 {
		return nil, errors.New("missing input resource kind")
	}

	// Use config from file
	if len(resourceCfg) > 0 {
		return common.ReadConfigFromFile(resourceCfg)
	}

	// Use default config
	return loadDefaultConfig(args)
}

func loadDefaultConfig(args []string) (*common.ResourceConfig, error) {
	switch args[0] {
	case "mongors":
		return common.DefaultMongoReplicaSetConfig(), nil
	default:
		return nil, errors.New("kind is not found")
	}
}
