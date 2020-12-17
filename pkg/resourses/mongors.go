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

package resourses

import (
	"bytes"
	"fmt"
	"github.com/Kaiser925/devctl/pkg/common"
	"log"
	"os"
	"os/exec"
)

type MongoReplicaSet struct {
	Host    string
	DataDir string

	setupDir string
	kind     string
}

func NewMongoReplicaSet(config *common.ResourceConfig) Resource {
	return &MongoReplicaSet{
		DataDir:  config.DataDir,
		Host:     config.Host,
		setupDir: config.SetupDir,
		kind:     "MongoRS",
	}
}

func (m *MongoReplicaSet) Kind() string {
	return m.kind
}

func (m *MongoReplicaSet) Create() error {
	err := m.prepareFiles()
	if err != nil {
		return err
	}

	var outB, errB bytes.Buffer
	cmdStr := "docker-compose up --build -d"
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Dir = m.setupDir
	cmd.Stdout = &outB
	cmd.Stderr = &errB
	err = cmd.Run()
	if err != nil {
		log.Println(errB.String())
		log.Fatal(err.Error())
	}
	log.Println(outB.String())
	return nil
}

func (m *MongoReplicaSet) Delete() error {
	log.Println("delete mongo replicaset and data")

	cmd := exec.Command("docker-compose", "down")
	cmd.Dir = m.setupDir
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	err := os.RemoveAll(m.DataDir)
	if err != nil {
		return err
	}

	log.Println("mongo replica set has been removed")
	return nil
}

func (m *MongoReplicaSet) prepareFiles() error {
	err := os.Mkdir(m.setupDir, os.ModePerm)

	if err != nil && !os.IsExist(err) {
		return err
	}

	_, err = common.WriteFile(fmt.Sprintf("%s/setup.sh", m.setupDir), common.SETUP_SHELL)
	if err != nil {
		return err
	}
	_, err = common.WriteFile(fmt.Sprintf("%s/Dockerfile", m.setupDir), common.SETUP_DOCKER)
	if err != nil {
		return err
	}

	configFile := fmt.Sprintf("%s/replicaSet.js", m.setupDir)
	err = common.RenderTemplateFile(common.REPLICA_SET_CONFG, m, configFile)

	composeFile := fmt.Sprintf("%s/docker-compose.yaml", m.setupDir)
	err = common.RenderTemplateFile(common.MONGO_RS_DOCKER_COMPOSE, m, composeFile)

	return nil
}
