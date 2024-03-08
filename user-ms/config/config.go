package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type conf struct {
	DbDriver     string `mapstructure:"database.mongodb"`
	DbHost       string `mapstructure:"database.mongodb.host"`
	DbPort       string `mapstructure:"database.mongodb.port"`
	DbName       string `mapstructure:"database.mongodb.database"`
	DbUser       string `mapstructure:"database.mongodb.user"`
	DbPassword   string `mapstructure:"database.mongodb.password"`
	DbUseSSL     string `mapstructure:"database.mongodb.ssl"`
	DbAuthSource string `mapstructure:"database.mongodb.authsource"`
}

func LoadConfig(path string, filename string) (*conf, error) {
	var config conf
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	fmt.Println(config.DbHost)

	return &config, nil
}
