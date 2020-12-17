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
	"fmt"
	"github.com/Kaiser925/devctl/pkg/common"
	"strings"
)

type ResourceNotFoundErr struct {
	kind string
}

func (e *ResourceNotFoundErr) Error() string {
	return fmt.Sprintf("doesn't has resource type \"%s\"", e.kind)
}

// ResourceOperator is used to execute command.
type ResourceOperator struct {
	resources map[string]ResourceConstructor
}

// NewResourceOperator returns new ResourceOperator pointer.
func NewResourceOperator() *ResourceOperator {
	op := &ResourceOperator{
		resources: make(map[string]ResourceConstructor),
	}
	op.AddResource("MongoRS", NewMongoReplicaSet)
	return op
}

// AddResources add a batch kinds of Resource to operator.
func (r *ResourceOperator) AddResources(resources map[string]ResourceConstructor) {
	for kind, constructor := range resources {
		r.resources[strings.ToLower(kind)] = constructor
	}
}

// AddResource add single kind Resource to operator.
func (r *ResourceOperator) AddResource(kind string, constructor ResourceConstructor) {
	r.resources[strings.ToLower(kind)] = constructor
}

// CreateResource creates new Resource.
func (r *ResourceOperator) CreateResource(config *common.ResourceConfig) error {
	resource, err := r.constructResource(config)
	if err != nil {
		return err
	}
	return resource.Create()
}

// DeleteResource deletes specify Resource.
func (r *ResourceOperator) DeleteResource(config *common.ResourceConfig) error {
	resource, err := r.constructResource(config)
	if err != nil {
		return err
	}
	return resource.Delete()
}

// constructResource constructs Resource corresponding to config.
func (r *ResourceOperator) constructResource(config *common.ResourceConfig) (Resource, error) {
	constructor, ok := r.resources[strings.ToLower(config.Kind)]
	if !ok {
		return nil, &ResourceNotFoundErr{kind: config.Kind}
	}
	return constructor(config), nil
}
