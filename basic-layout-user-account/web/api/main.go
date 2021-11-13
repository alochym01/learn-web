package main

import (
	"database/sql"
	"fmt"

	"github.com/alochym01/learn-web/basic-layout-user-account/router"
)

func main() {
	fmt.Println("Server is starting with port 8080")
	DB, err1 := sql.Open("sqlite3", "alochym.db")

	if err1 != nil {
		panic(err1)
	}

	defer DB.Close()
	// logger := logger.NewLogger()
	router := router.Router(DB)
	router.Run("localhost:8080")
}
