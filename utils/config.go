package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// LoadConfig config file
func LoadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Sprintf("Error on load config: %s \n", err.Error()))
	}
}
