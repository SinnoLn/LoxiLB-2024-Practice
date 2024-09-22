/*
 * Copyright (c) 2022 NetLOX Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package create

import (
	"fmt"
	"github.com/spf13/cobra"
)

//func CreateCmd(restOptions *api.RESTOptions) *cobra.Command {
func CreateCmd() *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a Load balance features in the LoxiLB.",
		Long: `Create a Load balance features in the LoxiLB.
Create - Service type external load-balancer, Vlan, Vxlan, Qos Policies, 
	 Endpoint client,FDB, IPaddress, Neighbor, Route,Firewall, Mirror, Session, UlCl
`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Create called!\n", args)
		},
	}

	//createCmd.AddCommand(NewCreateLoadBalancerCmd(restOptions))
	return createCmd
}
