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
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		userDb := database.NewUserDB(Db, "user")
		find := usecases.NewFindUserUseCase(userDb)
		user, err := find.Execute(ctx, usecases.FindUserInput{Id: userId, Email: email})
		if err != nil {
			return err
		}

		deactivate := usecases.NewDeleteUserUseCase(userDb)
		err = deactivate.Execute(ctx, usecases.DeleteUserInput{UserId: user.UserId})
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
