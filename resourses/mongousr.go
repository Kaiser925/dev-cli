package resourses

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDB struct {
	host string
}

func NewMongoDB(host string) *MongoDB {
	return &MongoDB{
		host: host,
	}
}

func (m *MongoDB) Create(db, user, passwd string) error {
	log.Println(fmt.Sprintf("create %s:%s for %s on %s", user, passwd, db, m.host))
	uri := fmt.Sprintf("mongodb://%s:27010,%s:27011,%s:27012/?replicaSet=rs0", m.host, m.host, m.host)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	if err := client.Connect(ctx); err != nil {
		return err
	}

	r := client.Database(db).RunCommand(
		context.Background(),
		bson.D{
			{"createUser", user},
			{"pwd", passwd},
			{"roles", []bson.M{{"role": "readWrite", "db": db}}},
		})

	if r.Err() != nil {
		return r.Err()
	}

	log.Println(fmt.Sprintf("create done,use \"mongodb://%s:%s@%s:27010,%s:27011,%s:27012/%s?replicaSet=rs0\" to connect",
		user, passwd, m.host, m.host, m.host, db))
	return nil
}

func (m *MongoDB) Remove(db, user string) error {
	return nil
}
