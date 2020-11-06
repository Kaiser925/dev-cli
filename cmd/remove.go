package cmd

import (
	"errors"
	"fmt"
	"github.com/Kaiser925/devctl/common"
	"github.com/Kaiser925/devctl/resourses"
	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   "delete <resource>",
	Short: "Delete local resource",
	Args:  cobra.MinimumNArgs(1),
	Long: `Delete local resource, such as local mongo replica set.

# remove mongo replica set, and delete local data.
# Note need root.
devctl delete mongors
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		remover := &remover{}
		return remover.remove(args[0], args[1:])
	},
}

type remover struct{}

func (r *remover) remove(res string, param []string) error {
	switch res {
	case "mongors":
		return r.removeMongoRS(param)
	default:
		return errors.New(fmt.Sprintf("no resource named %s", res))
	}
}

func (r *remover) removeMongoRS(param []string) error {
	config := common.DefaultMongoReplicaSetConfig()
	if len(param) == 1 {
		config.DataDir = param[0]
	}
	mongors := resourses.NewMongoReplicaSet(config)
	return mongors.Delete()
}
