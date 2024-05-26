package database

import (
	"fmt"
	"goapirest/configs"

	"github.com/jackc/pgx"
)


func DBConnection () (*pgx.Conn, error) {
	config, err := configs.LoadEnv(".env")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	postgresConfig := pgx.ConnConfig{
		Host: config.DBHost,
		Port: config.DBPort,
		Database: config.DBName,
		User: config.DBUser,
		Password: config.DBPassword,
	}

	conn, err := pgx.Connect(postgresConfig)
	if err != nil {
		panic(fmt.Errorf("Unable to connect to database: %v", err))
	}
	fmt.Print(conn.PID())

	return conn, nil

}
