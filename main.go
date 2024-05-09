package main

import (
	"fmt"
	"goapirest/configs"
)

func init() {
	_, err := configs.LoadEnv(".env")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {

}
