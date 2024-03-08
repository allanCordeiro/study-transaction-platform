/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/database"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/usecases"
	"github.com/spf13/cobra"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

var (
	name     string
	email    string
	userType string
	password string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new user",
	Long:  "create a new user",
	RunE:  runCreateUser(),
}

func runCreateUser() RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		userDb := database.NewUserDB(Db)
		create := usecases.NewCreateUserUseCase(userDb)
		_, err := create.Execute(context.TODO(), usecases.CreateUserInput{
			Name:     name,
			Email:    email,
			Password: password,
			UserType: userType,
		})
		if err != nil {
			return err
		}

		return nil
	}
}

func init() {

	userCmd.AddCommand(createCmd)

	//flags
	createCmd.Flags().StringVarP(&name, "name", "n", "", "User's name")
	createCmd.Flags().StringVarP(&email, "email", "e", "", "User's email address")
	createCmd.Flags().StringVarP(&userType, "type", "t", "", "User's type(customer or distributor)")
	createCmd.Flags().StringVarP(&password, "password", "p", "", "User's password")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("email")
	createCmd.MarkFlagRequired("type")
	createCmd.MarkFlagRequired("password")
}
