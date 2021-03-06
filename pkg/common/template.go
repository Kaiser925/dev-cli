/*
 * Developed by Kaiser925 on 2020/12/17.
 * Lasted modified 2020/11/3.
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

const REPLICA_SET_CONFG = `
rs.initiate({
  _id: "rs0",
  members: [
    { _id: 0, host: "{{ .Host }}:27010" },
    { _id: 1, host: "{{ .Host }}:27011" },
    { _id: 2, host: "{{ .Host }}:27012" },
  ],
});

rs.conf();
`

const SETUP_SHELL = `#!/usr/bin/env bash

set -e

sleep 10 | echo Sleeping

mongo mongodb://mongo-rs0-0:27017 < replicaSet.js`

const SETUP_DOCKER = `FROM mongo:4

COPY replicaSet.js .
COPY setup.sh .
`

const MONGO_RS_DOCKER_COMPOSE = `
version: "3"
services:
  mongo-rs0-0:
    container_name: m-0
    image: mongo:4
    expose:
      - 27017
    ports:
      - 27010:27017
    volumes:
      - {{ .DataDir }}/rs0-0:/data/db
    restart: always
    entrypoint: ["/usr/bin/mongod", "--bind_ip_all", "--replSet", "rs0"]
    depends_on:
      - "mongo-rs0-1"
      - "mongo-rs0-2"

  mongo-rs0-1:
    container_name: m-1
    image: mongo:4
    expose:
      - 27017
    ports:
      - 27011:27017
    volumes:
      - {{ .DataDir }}/rs0-1:/data/db
    restart: always
    entrypoint: ["/usr/bin/mongod", "--bind_ip_all", "--replSet", "rs0"]

  mongo-rs0-2:
    container_name: m-2
    image: mongo:4
    expose:
      - 27017
    ports:
      - 27012:27017
    volumes:
      - {{ .DataDir }}/rs0-2:/data/db
    restart: always
    entrypoint: ["/usr/bin/mongod", "--bind_ip_all", "--replSet", "rs0"]

  setup-rs:
    image: "setup-rs"
    build: .
    entrypoint: ["/bin/bash", "./setup.sh"]
    depends_on:
      - "mongo-rs0-0"

  adminmongo:
    image: "mrvautin/adminmongo"
    container_name: m-ui
    environment:
      - HOST=0.0.0.0
    ports:
      - "1234:1234"
    restart: always
`
