package common

type MongoReplicaSetConfig struct {
	Host     string `yaml:"Host"`
	DataDir  string `yaml:"DataDir"`
	SetupDir string `yaml:"SetupDir"`
}

type MongoDBConfig struct {
	Host         string `yaml:"Host"`
	DataBaseName string `yaml:"DataBaseName"`
	User         string `yaml:"User"`
	Password     string `yaml:"Password"`
}

func DefaultMongoReplicaSetConfig() *MongoReplicaSetConfig {
	host, _ := GetLocalIP()
	return &MongoReplicaSetConfig{
		Host:     host,
		DataDir:  "/mnt/data/mongo",
		SetupDir: "./.mongo-setup",
	}
}

func DefaultMongoDBConfig() *MongoDBConfig {
	host, _ := GetLocalIP()
	return &MongoDBConfig{
		Host:         host,
		DataBaseName: "MongoDB",
		User:         "admin",
		Password:     "admin",
	}
}
