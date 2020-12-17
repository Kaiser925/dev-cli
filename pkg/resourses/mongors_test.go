package resourses

import (
	"github.com/Kaiser925/devctl/pkg/common"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewMongoReplicaSet(t *testing.T) {
	config := new(common.ResourceConfig)
	mongors := NewMongoReplicaSet(config)
	assert.Equal(t, mongors.Kind(), "MongoRS")
}
