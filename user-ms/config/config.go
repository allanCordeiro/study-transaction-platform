package config

import (
	"github.com/spf13/viper"
)

type conf struct {
	DbDriver     string `mapstructure:"driver"`
	DbHost       string `mapstructure:"host"`
	DbPort       string `mapstructure:"port"`
	DbName       string `mapstructure:"database"`
	DbUser       string `mapstructure:"user"`
	DbPassword   string `mapstructure:"password"`
	DbUseSSL     bool   `mapstructure:"ssl"`
	DbAuthSource string `mapstructure:"authSource"`
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

	err = viper.UnmarshalKey("database", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
