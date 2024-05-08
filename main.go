package main

import (
	"fmt"
	"goapirest/configs"
)

func main() {
	config, err := configs.LoadEnv(".env")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println(*config)
}
