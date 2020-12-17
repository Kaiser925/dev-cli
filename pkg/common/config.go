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
