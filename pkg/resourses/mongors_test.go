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
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewMongoReplicaSet(t *testing.T) {
	config := new(common.ResourceConfig)
	mongors := NewMongoReplicaSet(config)
	assert.Equal(t, mongors.Kind(), "MongoRS")
}
