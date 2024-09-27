package create

import (
	"clitest/api"
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

			// 성공 메시지를 출력 (데이터베이스 작업은 생략)
			fmt.Printf("Account created successfully (simulated):\n")
			fmt.Printf("UserID: %s, Password: %s, Email: %s\n", userID, password, email)
		},
	}

	return CreateAccountCmd
}
