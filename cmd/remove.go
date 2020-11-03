package cmd

import (
	"errors"
	"fmt"
	"github.com/Kaiser925/devctl/common"
	"github.com/Kaiser925/devctl/resourses"
	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove <resource>",
	Short: "Remove local resource",
	Args:  cobra.MinimumNArgs(1),
	Long: `Remove local resource, such as local mongo replica set.

# remove mongo replica set, and delete local data.
# Note need root.
devctl remove mongors
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
	dataDir := "./data"
	if len(param) == 1 {
		dataDir = param[0]
	}
	ip, err := common.GetLocalIP()
	if err != nil {
		return err
	}
	mongors := resourses.NewMongoReplicaSet(dataDir, ip)
	return mongors.Remove()
}
