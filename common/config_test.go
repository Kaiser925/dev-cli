package common

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDefaultMongoReplicaSetConfig(t *testing.T) {
	config := DefaultMongoReplicaSetConfig()

	assert.Equal(t, config.Kind, "MongoReplicaSet")
	assert.Equal(t, config.DataDir, "/mnt/data/mongo")
	assert.Equal(t, config.SetupDir, "./.devctl-setup")
}

func TestDefaultMongoDBConfig(t *testing.T) {
	config := DefaultMongoDBConfig()

	assert.Equal(t, config.Kind, "MongoDB")
	assert.Equal(t, config.DatabaseName, "MongoDB")
	assert.Equal(t, config.Password, "admin")
	assert.Equal(t, config.User, "admin")
}
