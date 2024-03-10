/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/database"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/usecases"
	"github.com/spf13/cobra"
)

// deactivateCmd represents the deactivate command
var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactive user account",
	Long:  "Deactive user account",
	RunE:  runDeactivateUser(),
}

func runDeactivateUser() RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		userDb := database.NewUserDB(Db)
		deactivate := usecases.NewDeleteUserUseCase(userDb)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err := deactivate.Execute(ctx, usecases.DeleteUserInput{UserId: userId})
		if err != nil {
			return err
		}

		fmt.Println("user deactivated")
		return nil
	}
}

func init() {
	userCmd.AddCommand(deactivateCmd)

	deactivateCmd.Flags().StringVarP(&email, "email", "e", "", "User's email address")
	deactivateCmd.Flags().StringVarP(&userId, "id", "i", "", "User's id")
	deactivateCmd.MarkFlagsOneRequired("email", "id")
}
