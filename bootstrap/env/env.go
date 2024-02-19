package env

import (
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ApiKey        string `mapstructure:"API_KEY"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	ATSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	ClientURL     string `mapstructure:"CLIENT_URL"`
}

var ProcEnv = GetEnv()

func GetEnv() *Env {
	env := &Env{}

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		logger.FatalLog(layers.Env, "cant find the env file", err)
	}

	if err := viper.Unmarshal(env); err != nil {
		logger.FatalLog(layers.Env, "env variables cant be loaded", err)
	}

	if env.AppEnv == "development" {
		logger.Logger(layers.Env, "server application is running on development mode")
	}

	return env
}
