package delete

import (
	"clitest/api"
	"fmt"

	"github.com/spf13/cobra"
)

// func DeleteCmd(restOptions *api.RESTOptions) *cobra.Command {
func DeleteCmd(restOptions *api.RESTOptions) *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a Load balance features in the LoxiLB.",
		Long: `Delete a Load balance features in the LoxiLB.
Delete - Service type external load-balancer, Vlan, Vxlan, Qos Policies,
	 Endpoint client,FDB, IPaddress, Neighbor, Route,Firewall, Mirror, Session, UlCl',

`, Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Delete called!\n", args)
		},
	}

	//'delete account' 명령어 추가
	deleteCmd.AddCommand(NewDeleteAccountCmd(restOptions))
	return deleteCmd
}
