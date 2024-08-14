package main

import (
	"fmt"

	"github.com/spf13/viper"
	"quiz/internal/cli"
)

func main() {
	viper.SetConfigFile("./.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.AutomaticEnv()

	cli.Execute()
}
