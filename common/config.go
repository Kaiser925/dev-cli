package common

type ResourceConfig struct {
	Kind         string `yaml:"kind"`
	Host         string `yaml:"host"`
	DataDir      string `yaml:"data-dir"`
	SetupDir     string `yaml:"setup-dir"`
	DatabaseName string `yaml:"database"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
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
