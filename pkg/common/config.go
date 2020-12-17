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

package common

import (
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type ResourceConfig struct {
	Kind         string `yaml:"kind"`
	Host         string `yaml:"host"`
	DataDir      string `yaml:"data-dir"`
	SetupDir     string `yaml:"setup-dir"`
	DatabaseName string `yaml:"database"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
}

func NewResourceConfig() *ResourceConfig {
	return new(ResourceConfig)
}

func ReadConfigFromFile(name string) (*ResourceConfig, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return ReadConfig(file)
}

func ReadConfig(reader io.Reader) (*ResourceConfig, error) {
	config := &ResourceConfig{}
	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}

func DefaultMongoReplicaSetConfig() *ResourceConfig {
	host, _ := GetLocalIP()
	return &ResourceConfig{
		Kind:     "MongoReplicaSet",
		Host:     host,
		DataDir:  "/mnt/data/mongo",
		SetupDir: "./.devctl-setup",
	}
}

func DefaultMongoDBConfig() *ResourceConfig {
	host, _ := GetLocalIP()
	return &ResourceConfig{
		Kind:         "MongoDB",
		Host:         host,
		DatabaseName: "MongoDB",
		User:         "admin",
		Password:     "admin",
	}
}
