package delete

import (
	"clitest/api"
	"fmt"

	"github.com/spf13/cobra"
)

// NewDeleteAccountCmd 함수는 'delete account' 명령어를 정의하고 처리하는 함수
func NewDeleteAccountCmd(restOptions *api.RESTOptions) *cobra.Command {
	var DeleteAccountCmd = &cobra.Command{
		Use:   "account [user_id]", // CLI 명령어로 user_id 인자를 받아 삭제할 대상 지정
		Short: "Delete an account by user_id",
		Long:  `Delete a specific account by providing the user_id.`,
		Args:  cobra.ExactArgs(1), // user_id 인자를 1개 받아야 함
		Run: func(cmd *cobra.Command, args []string) {
			// user_id 가져오기
			userID := args[0]

			// 실제로 데이터베이스 작업이 없으므로 단순히 성공 메시지 출력
			fmt.Printf("Account %s deleted successfully (simulated)\n", userID)
		},
	}

	return DeleteAccountCmd
}
