/*
 * Developed by Kaiser925 on 2020/12/17.
 * Lasted modified 2020/11/14.
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

package common

import (
	"github.com/magiconair/properties/assert"
	"strings"
	"testing"
)

var configStr = `
kind: test
host: localhost
data-dir: ./data
setup-dir: ./setup
`

func TestReadConfig(t *testing.T) {
	reader := strings.NewReader(configStr)
	config, err := ReadConfig(reader)

	assert.Equal(t, err, nil)
	assert.Equal(t, config.Kind, "test")
	assert.Equal(t, config.Host, "localhost")
	assert.Equal(t, config.DataDir, "./data")
	assert.Equal(t, config.SetupDir, "./setup")
	assert.Equal(t, config.User, "")
}

func TestReadConfigFromFile(t *testing.T) {
	config, err := ReadConfigFromFile("tmp-config.yaml")
	assert.Equal(t, err, nil)
	assert.Equal(t, config.Kind, "test")
	assert.Equal(t, config.DatabaseName, "db")
	assert.Equal(t, config.User, "usr")
	assert.Equal(t, config.Password, "pass")
	assert.Equal(t, config.Kind, "test")
	assert.Equal(t, config.Host, "192.168.1.1")
	assert.Equal(t, config.DataDir, "./data")
	assert.Equal(t, config.SetupDir, "./setup")
}

func TestDefaultMongoReplicaSetConfig(t *testing.T) {
	config := DefaultMongoReplicaSetConfig()

	assert.Equal(t, config.Kind, "MongoReplicaSet")
	assert.Equal(t, config.DataDir, "/mnt/data/mongo")
	assert.Equal(t, config.SetupDir, "./.dev-cli-setup")
}
