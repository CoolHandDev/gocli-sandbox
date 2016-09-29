package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	fmt.Println("username command")
	RootCmd.AddCommand(usernameCmd)
}

var usernameCmd = &cobra.Command{
	Use:   "username",
	Short: "SQL Server user name",
	Long:  `Enter the user name for SQL Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is the username command")
	},
}
