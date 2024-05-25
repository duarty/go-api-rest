package database

import (
	"fmt"
	"goapirest/configs"
	"os"

	"github.com/jackc/pgx"
)


func DBConnection () {
	config, _ := configs.LoadEnv(".env")
	postgresConfig := pgx.ConnConfig{
		Host: config.DBHost,
		Port: config.DBPort,
		Database: config.DBName,
		User: config.DBUser,
		Password: config.DBPassword,
	}
	conn, err := pgx.Connect(postgresConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// ctx := context.Background()
	// ping := conn.Ping(ctx)

	// select {
	// case <-time.After(3 * time.Second):
	// 	ping.Error()
	// case <-ctx.Done():
	// 	fmt.Println("Conexão com o banco de dados realizada com sucesso!")
	// }

	// defer conn.Close()
	fmt.Print(conn.PID())
	// var name string
	// var weight int64
	// err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(name, weight)
}
