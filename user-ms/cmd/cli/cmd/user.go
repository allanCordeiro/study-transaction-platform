/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/config"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Set of commands to manage user's database",
	Long:  "Set of commands to manage user's database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("user called")
	},
}

func init() {
	config, err := config.LoadConfig(".", "configs")
	if err != nil {
		panic(err)
	}
	fmt.Println(config.DbHost)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	uri := fmt.Sprintf("%s://%s:%s@%s:%s/%s?ssl=%s&authSource=%s",
		config.DbDriver,
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
		config.DbUseSSL,
		config.DbAuthSource)
	Db, err = GetDatabase(ctx, uri)
	if err != nil {
		panic(err)
	}
	defer Db.Client().Disconnect(ctx)

	rootCmd.AddCommand(userCmd)
}

func GetDatabase(ctx context.Context, dbUri string) (*mongo.Database, error) {
	fmt.Println(dbUri)
	//mongoURI := "mongodb://root:root@localhost:27017/teste?ssl=false&authSource=admin"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		return nil, err
	}

	return client.Database("user"), nil
}
