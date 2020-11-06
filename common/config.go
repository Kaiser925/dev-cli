package common

type MongoReplicaSetConfig struct {
	Host     string `yaml:"Host"`
	DataDir  string `yaml:"DataDir"`
	SetupDir string `yaml:"SetupDir"`
}

func DefaultMongoReplicaSetConfig() *MongoReplicaSetConfig {
	host, _ := GetLocalIP()
	return &MongoReplicaSetConfig{
		Host:     host,
		DataDir:  "/mnt/data/mongo",
		SetupDir: "./.mongo-setup",
	}
}
