package config

import (
	"os"
    "github.com/spf13/viper"
)

func GetEnv(key, defaultVal string) string {

    viper.SetConfigFile(".env")
    viper.ReadInConfig()

    if value := viper.GetString(key); value != "" {
        return value
    }

    if value, exists := os.LookupEnv(key); exists {
        return value
    }

    return defaultVal
}
