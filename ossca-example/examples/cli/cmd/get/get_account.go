package get

import (
	"clitest/api"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func NewGetAccountCmd(restOptions *api.RESTOptions) *cobra.Command {
	var GetAccountCmd = &cobra.Command{
		Use:   "account",
		Short: "Get a account",
		Long:  `It shows account informations.`,
		Run: func(cmd *cobra.Command, args []string) {
			client := api.NewLoxiClient(restOptions)
			ctx := context.TODO()
			var cancel context.CancelFunc
			if restOptions.Timeout > 0 {
				ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
				defer cancel()
			}
			resp, err := client.Account().Get(ctx)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			if resp.StatusCode == http.StatusOK {
				PrintGetAccountResult(resp, *restOptions)
				return
			}

		},
	}

	return GetAccountCmd
}

func PrintGetAccountResult(resp *http.Response, o api.RESTOptions) {
	AccountResp := api.AccountModGet{}
	var data [][]string
	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
		return
	}

	if err := json.Unmarshal(resultByte, &AccountResp); err != nil {
		fmt.Printf("Error: Failed to unmarshal HTTP response: (%s)\n", err.Error())
		return
	}

	resultIndent, _ := json.MarshalIndent(AccountResp, "", "    ")
	fmt.Println(string(resultIndent))

	// Table Init
	table := TableInit()

	// Making fdb data
	for _, account := range AccountResp.Attr {

		table.SetHeader([]string{"user id", "password", "email"})
		data = append(data, []string{account.UserID, account.Password, account.Email})

	}
	// Rendering the fdb data to table
	TableShow(data, table)
}
