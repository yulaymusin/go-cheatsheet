package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Open up our database connection.
	conn, _ := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/db_name")

	// defer the close till after the main function has finished
	// executing
	defer conn.Close(context.Background())
	var greeting string
	//
	conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	fmt.Println(greeting)
}
