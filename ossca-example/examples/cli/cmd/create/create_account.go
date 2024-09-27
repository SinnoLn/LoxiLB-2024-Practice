package create

import (
	"clitest/api"
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// NewCreateAccountCmd 함수는 'create account' 명령어를 정의하고 처리하는 함수
func NewCreateAccountCmd(restOptions *api.RESTOptions) *cobra.Command {
	var CreateAccountCmd = &cobra.Command{
		Use:   "account [user_id] [password] [email]", // user_id, password, email 인자를 받음
		Short: "Create a new account by user_id, password, and email",
		Long:  `Create a new account by providing user_id, password, and email.`,
		Args:  cobra.ExactArgs(3), // 인자 3개가 필요함
		Run: func(cmd *cobra.Command, args []string) {
			// 인자로 받은 계정 정보 가져오기
			userID := args[0]
			password := args[1]
			email := args[2]

			// api simulation
			client := api.NewLoxiClient(restOptions) // REST 클라이언트 생성
			ctx := context.TODO()                    // 컨텍스트 생성

			newAccount := api.AccountGet{
				UserID:   userID,
				Password: password,
				Email:    email,
			}

			resp, err := client.Account().Create(ctx, newAccount) // 계정 생성 요청
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}

			jsonFlag, _ := cmd.Flags().GetBool("json") // json 플래그 확인
			wideFlag, _ := cmd.Flags().GetBool("wide") // wide 플래그 확인

			if jsonFlag {
				// json 출력
				accountData, _ := json.MarshalIndent(newAccount, "", "    ")
				fmt.Println(string(accountData))
				return // json 출력 후 종료
			}

			if wideFlag {
				// wide 출력
				fmt.Printf("Account created successfully (wide):\n")
				fmt.Printf("UserID: %s, Password: %s, Email %s, Status: %d\n", userID, password, email, resp.StatusCode)
				return // wide 출력 후 종료
			}

			// 성공 메시지를 출력 (데이터베이스 작업은 생략)
			fmt.Printf("Account created successfully (simulated):\n")
			fmt.Printf("UserID: %s, Password: %s, Email: %s\n", userID, password, email)
		},
	}

	// json 및 wide 플래그 추가
	CreateAccountCmd.Flags().Bool("json", false, "Output result in JSON format")
	CreateAccountCmd.Flags().Bool("wide", false, "Output result in wide format")

	return CreateAccountCmd
}
