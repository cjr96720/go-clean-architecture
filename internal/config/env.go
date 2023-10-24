package config

import (
	"github.com/spf13/viper"

	"go-clean-architecture/internal/utils"
)

type Env struct {
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresName     string `mapstructure:"POSTGRES_DB"`
	APIAddress       string `mapstructure:"API_ADDRESS"`
}

func NewEnv() *Env {
	env := Env{}

	// `SetConfigFile` defines the path, name, extension of the configuration file
	viper.SetConfigFile(".env")

	// `ReadInConfig` will discover and load the configuration file
	err := viper.ReadInConfig()
	utils.ErrorPanic(err)

	// `Unmarshal` unmarshals the config into struct
	err = viper.Unmarshal(&env)
	utils.ErrorPanic(err)

	return &env
}
