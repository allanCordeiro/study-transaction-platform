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

// reactivateCmd represents the reactivate command
var reactivateCmd = &cobra.Command{
	Use:   "reactivate",
	Short: "Reactivate user account",
	Long:  "Reactivate user account",
	RunE:  runReactivateUser(),
}

func runReactivateUser() RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		userDb := database.NewUserDB(Db)
		find := usecases.NewFindUserUseCase(userDb)
		user, err := find.Execute(ctx, usecases.FindUserInput{Id: userId, Email: email})
		if err != nil {
			return err
		}

		reactivate := usecases.NewUpdateUserUseCase(userDb)
		err = reactivate.Execute(ctx, usecases.UpdateUserUseCaseInput{
			UserId:   user.UserId,
			Name:     user.Name,
			Password: user.Password,
			UserType: user.UserType,
			IsActive: true,
		})
		if err != nil {
			return err
		}

		fmt.Println("user has been activated")
		return nil
	}
}

func init() {
	userCmd.AddCommand(reactivateCmd)

	reactivateCmd.Flags().StringVarP(&email, "email", "e", "", "User's email address")
	reactivateCmd.Flags().StringVarP(&userId, "id", "i", "", "User's id")
	reactivateCmd.MarkFlagsOneRequired("email", "id")
}
