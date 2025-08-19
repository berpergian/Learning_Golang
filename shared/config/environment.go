package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	RabbitHost             string `mapstructure:"RABBIT_HOST"`
	RabbitPort             string `mapstructure:"RABBIT_PORT"`
	RabbitUser             string `mapstructure:"RABBIT_USER"`
	RabbitPass             string `mapstructure:"RABBIT_PASS"`
	MessageExchange        string `mapstructure:"MESSAGE_EXCHANGE"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	Issuer                 string `mapstructure:"ISSUER"`
}

func ReadEnvironment() *Env {
	env := Env{}

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "..")

	envPath := filepath.Join(basePath, ".env")

	viper.SetConfigFile(envPath)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
