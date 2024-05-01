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

func init() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "."
	}
	viper.AddConfigPath(path)
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
	env := os.Getenv("ENVIRONMENT")
	if env != "" {
		return viper.GetString(env + "." + key)
	}
	return viper.GetString("default" + "." + key)
}
