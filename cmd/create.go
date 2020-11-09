package cmd

import (
	"errors"
	"fmt"
	"github.com/Kaiser925/devctl/common"
	"github.com/Kaiser925/devctl/resourses"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <resource>",
	Short: "create local resource",
	Args:  cobra.MinimumNArgs(1),
	Long: `create local resource, such as local mongo replica set.

# create new mongo replica set
devctl create mongors

# create new user password for mongo database. If database not exists, it will be created.
devctl create mongousr <dbname> <user> <passwd>
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		creater := newCreate()
		res := args[0]
		return creater.create(res, args[1:])
	},
}

type create struct{}

func newCreate() *create {
	return &create{}
}

func (c *create) create(res string, param []string) error {
	switch res {
	case "mongors":
		return c.createReplicaSet(param)
	case "mongousr":
		return c.createMongoUsr(param)
	default:
		return errors.New(fmt.Sprintf("no resource named %s", res))
	}
}

func (c *create) createReplicaSet(param []string) error {
	config := common.DefaultMongoReplicaSetConfig()
	if len(param) == 1 {
		config.DataDir = param[0]
	}
	mongors := resourses.NewMongoReplicaSet(config)
	return mongors.Create()
}

func (c *create) createMongoUsr(param []string) error {
	if len(param) < 3 {
		return errors.New("create database need database name, user and password")
	}
	config := common.DefaultMongoDBConfig()
	config.DatabaseName, config.User, config.Password = param[0], param[1], param[2]

	mongodb := resourses.NewMongoDB(config)

	return mongodb.Create()
}
