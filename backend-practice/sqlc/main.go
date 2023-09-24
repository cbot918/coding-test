package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

const dsn = "postgres://postgres:12345@localhost:5433/practice2?sslmode=disable"

func main() {
	db := NewDB(dsn)
	users, err := db.Q.ListUsers(context.Background())
	if err != nil {
		log.Fatal("db list users failed", err.Error())
	}
	fmt.Println("rows: ")
	logj(users)

}

func logj(v any) {
	res, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("marindent failed")
	}
	fmt.Println(string(res))
}
