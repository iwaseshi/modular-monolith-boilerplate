package config

import (
	"fmt"
	"modular-monolith-boilerplate/pkg/logger"
	"os"

	"github.com/spf13/viper"
)

var (
	isInitialized = false
)

func LoadServiceConfig(serviceName string) {
	viper.AddConfigPath(serviceName)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	isInitialized = true
}

func Get(key string) string {
	if !isInitialized {
		logger.Default().Warning("Config is not initialized")
	}
	str := ""
	env := os.Getenv("ENVIRONMENT")
	if env != "" {
		logger.Default().Info("ENVIRONMENT: " + env)
		return viper.GetString(env + "." + key)
	}
	str = viper.GetString("default" + "." + key)
	return str
}
