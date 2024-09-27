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

// NewGetAccountCmd 함수는 'get account' 명령어를 정의하고 처리하는 함수
func NewGetAccountCmd(restOptions *api.RESTOptions) *cobra.Command {
	var GetAccountCmd = &cobra.Command{
		Use:   "account",
		Short: "Get account details",
		Long:  `It shows account information.`,
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

			// json 및 wide 옵션 플래그 처리
			jsonFlag, _ := cmd.Flags().GetBool("json")
			wideFlag, _ := cmd.Flags().GetBool("wide")

			if resp.StatusCode == http.StatusOK {
				if jsonFlag {
					// JSON 형식으로 출력
					PrintGetAccountResultAsJSON(resp)
				} else if wideFlag {
					// wide 형식으로 출력
					PrintGetAccountResultAsWide(resp, *restOptions)
				} else {
					// 기본 출력
					PrintGetAccountResult(resp, *restOptions)
				}
				return
			}
		},
	}

	// json 및 wide 플래그 추가
	GetAccountCmd.Flags().Bool("json", false, "Output result in JSON format")
	GetAccountCmd.Flags().Bool("wide", false, "Output result in wide format")

	return GetAccountCmd
}

// PrintGetAccountResult 함수는 기본 테이블 형식으로 계정 정보를 출력
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

// PrintGetAccountResultAsJSON 함수는 계정 정보를 JSON 형식으로 출력
func PrintGetAccountResultAsJSON(resp *http.Response) {
	AccountResp := api.AccountModGet{}
	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
		return
	}

	if err := json.Unmarshal(resultByte, &AccountResp); err != nil {
		fmt.Printf("Error: Failed to unmarshal HTTP response: (%s)\n", err.Error())
		return
	}

	// JSON 형식으로 출력
	resultIndent, _ := json.MarshalIndent(AccountResp, "", "    ")
	fmt.Println(string(resultIndent))
}

// PrintGetAccountResultAsWide 함수는 wide 옵션에 맞게 확장된 형식으로 출력
func PrintGetAccountResultAsWide(resp *http.Response, o api.RESTOptions) {
	AccountResp := api.AccountModGet{}
	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
		return
	}

	if err := json.Unmarshal(resultByte, &AccountResp); err != nil {
		fmt.Printf("Error: Failed to unmarshal HTTP response: (%s)\n", err.Error())
		return
	}

	// Wide 형식으로 더 많은 정보 출력 (확장된 출력)
	for _, account := range AccountResp.Attr {
		fmt.Printf("Account Details (WIDE): UserID: %s, Password: %s, Email: %s\n", account.UserID, account.Password, account.Email)
	}
}
