package resourses

import (
	"context"
	"fmt"
	"github.com/Kaiser925/devctl/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDB struct {
	config *common.MongoDBConfig
}

func NewMongoDB(config *common.MongoDBConfig) *MongoDB {
	return &MongoDB{
		config,
	}
}

func (m *MongoDB) Create() error {
	uri := fmt.Sprintf("mongodb://%s:27010,%s:27011,%s:27012/?replicaSet=rs0",
		m.config.Host, m.config.Host, m.config.Host)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	if err := client.Connect(ctx); err != nil {
		return err
	}

	r := client.Database(m.config.DataBaseName).RunCommand(
		context.Background(),
		bson.D{
			{"createUser", m.config.User},
			{"pwd", m.config.Password},
			{"roles", []bson.M{{"role": "readWrite", "db": m.config.DataBaseName}}},
		})

	if r.Err() != nil {
		return r.Err()
	}

	dbURI := fmt.Sprintf("mongodb://%s:%s@%s:27010,%s:27011,%s:27012/%s?replicaSet=rs0",
		m.config.User,
		m.config.Password,
		m.config.Host,
		m.config.Host,
		m.config.Host,
		m.config.DataBaseName)

	log.Println(fmt.Sprintf("create done,use \"%s\" to connect", dbURI))
	return nil
}

func (m *MongoDB) Remove() error {
	return nil
}
