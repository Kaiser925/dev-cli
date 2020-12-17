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
	assert.Equal(t, config.SetupDir, "./.devctl-setup")
}

func TestDefaultMongoDBConfig(t *testing.T) {
	config := DefaultMongoDBConfig()

	assert.Equal(t, config.Kind, "MongoDB")
	assert.Equal(t, config.DatabaseName, "MongoDB")
	assert.Equal(t, config.Password, "admin")
	assert.Equal(t, config.User, "admin")
}
