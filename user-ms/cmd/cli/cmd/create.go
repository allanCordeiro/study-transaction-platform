/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/database"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/usecases"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		mongoURI := "mongodb://root:root@localhost:27017/teste?ssl=false&authSource=admin"
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
		if err != nil {
			panic(err)
		}
		defer client.Disconnect(context.TODO())

		db := database.NewUserDB(client.Database("user"))
		create := usecases.NewCreateUserUseCase(db)
		_, err = create.Execute(context.TODO(), usecases.CreateUserInput{
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVarP(&name, "name", "n", "", "User's name")
	createCmd.Flags().StringVarP(&email, "email", "e", "", "User's email address")
	createCmd.Flags().StringVarP(&userType, "type", "t", "", "User's type(customer or distributor)")
	createCmd.Flags().StringVarP(&password, "password", "p", "", "User's password")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("email")
	createCmd.MarkFlagRequired("type")
	createCmd.MarkFlagRequired("password")
}
