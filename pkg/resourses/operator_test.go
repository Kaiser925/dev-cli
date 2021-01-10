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
	"github.com/Kaiser925/dev-cli/pkg/common"
	"github.com/Kaiser925/dev-cli/pkg/resourses/mocks"
	"github.com/magiconair/properties/assert"
	"testing"
)

var (
	mockKind        = "MockResource"
	mockResource    = &mocks.Resource{}
	mockConstructor = func(config *common.ResourceConfig) Resource { return mockResource }
)

func TestNewResourceOperator(t *testing.T) {
	op := NewResourceOperator()
	assert.Equal(t, len(op.resources), 1)
}

func TestResourceOperator_AddResources(t *testing.T) {
	resources := map[string]ResourceConstructor{
		"A": func(config *common.ResourceConfig) Resource { return nil },
		"B": func(config *common.ResourceConfig) Resource { return nil },
	}

	op := &ResourceOperator{
		resources: make(map[string]ResourceConstructor),
	}

	op.AddResources(resources)

	assert.Equal(t, len(op.resources), len(resources))
}

func TestResourceOperator_ConstructResource(t *testing.T) {
	mockResource.On("Kind").Return(mockKind)

	op := &ResourceOperator{
		resources: make(map[string]ResourceConstructor),
	}

	op.AddResource(mockResource.Kind(), mockConstructor)

	config := &common.ResourceConfig{Kind: mockKind}
	resource, err := op.constructResource(config)

	assert.Equal(t, err, nil)
	assert.Equal(t, resource.Kind(), config.Kind)
}

func TestResourceOperator_CreateResource(t *testing.T) {
	mockResource.On("Create").Return(nil)
	op := NewResourceOperator()
	op.AddResource(mockKind, mockConstructor)

	config := &common.ResourceConfig{Kind: mockKind}

	err := op.CreateResource(config)
	assert.Equal(t, err, nil)

	config.Kind = "undefined"
	err = op.CreateResource(config)
	assert.Equal(t, err != nil, true)
}

func TestResourceOperator_DeleteResource(t *testing.T) {
	mockResource.On("Delete").Return(nil)
	op := NewResourceOperator()
	op.AddResource(mockKind, mockConstructor)

	config := &common.ResourceConfig{Kind: mockKind}

	err := op.DeleteResource(config)
	assert.Equal(t, err, nil)

	config.Kind = "undefined"
	err = op.DeleteResource(config)
	assert.Equal(t, err != nil, true)
}
