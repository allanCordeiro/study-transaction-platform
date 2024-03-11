/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"time"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/database"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/usecases"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update user data",
	Long:  "Update user data",
	RunE:  runUpdateUser(),
}

func runUpdateUser() RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		userDb := database.NewUserDB(Db)
		update := usecases.NewUpdateUserUseCase(userDb)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err := update.Execute(ctx, usecases.UpdateUserUseCaseInput{
			UserId:   userId,
			Name:     name,
			UserType: userType,
			Password: password,
		})
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	userCmd.AddCommand(updateCmd)

	//flags
	updateCmd.Flags().StringVarP(&userId, "id", "i", "", "User's id")
	updateCmd.Flags().StringVarP(&name, "name", "n", "", "User's name")
	updateCmd.Flags().StringVarP(&userType, "type", "t", "", "User's type(customer or distributor)")
	updateCmd.Flags().StringVarP(&password, "password", "p", "", "User's password")

}
