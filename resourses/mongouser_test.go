package resourses

import (
	"github.com/Kaiser925/devctl/common"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewMongoUser(t *testing.T) {
	config := new(common.ResourceConfig)
	user := NewMongoUser(config)
	assert.Equal(t, user.Kind(), "MongoUser")
}
