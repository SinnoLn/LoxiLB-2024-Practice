package cmd

import (
	"clitest/api"
	"clitest/cmd/create"
	"clitest/cmd/get"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	restOptions := &api.RESTOptions{}

	rootCmd.PersistentFlags().Int16VarP(&restOptions.Timeout, "timeout", "t", 10, "Set timeout")
	rootCmd.PersistentFlags().StringVarP(&restOptions.Protocol, "protocol", "", "http", "Set API server http/https")
	rootCmd.PersistentFlags().StringVarP(&restOptions.PrintOption, "output", "o", "", "Set output layer (ex.) wide, json)")
	rootCmd.PersistentFlags().StringVarP(&restOptions.ServerIP, "apiserver", "s", "127.0.0.1", "Set API server IP address")
	rootCmd.PersistentFlags().Int16VarP(&restOptions.ServerPort, "port", "p", 11111, "Set API server port number")

	rootCmd.AddCommand(create.CreateCmd())
	rootCmd.AddCommand(get.GetCmd(restOptions))

}
