package cmd

import (
	"api/internal/common/app"
	"context"
	"github.com/spf13/cobra"
)

var email string
var password string

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Create admin user",
	Long:  `Create admin user`,
	Run: func(cmd *cobra.Command, args []string) {
		app := app.NewApplication(context.Background())
		app.AuthModule.Commands.CreateUser("admin", email, password)
	},
}

func init() {
	rootCmd.AddCommand(usersCmd)
	usersCmd.Flags().StringVarP(&email, "email", "e", "", "User email")
	usersCmd.Flags().StringVarP(&password, "password", "p", "", "User pass")
}
