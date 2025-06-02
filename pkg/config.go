package pkg

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/domaincheck")

	viper.SetDefault("workers", 5)
	viper.SetDefault("delay", 250)
	viper.SetDefault("api", "https://api.name.com/v4/domains:checkAvailability")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("ℹ️ No config file found; using flags/env vars only.")
	}
}
