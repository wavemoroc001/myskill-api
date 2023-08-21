package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"myskill-api/db"
)

type Config struct {
	Database db.Database `mapstructure:"db" validate:"required"`
}

func LoadEnv() Config {
	viper.SetConfigFile("config-dev.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	var v *validator.Validate
	v = validator.New()
	if err := v.Struct(&config); err != nil {
		panic(err)
	}
	return config
}
