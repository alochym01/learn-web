package main

import (
	"database/sql"

	"github.com/alochym01/web-w-gin/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err1 := sql.Open("sqlite3", "alochym.db")

	if err1 != nil {
		panic(err1)
	}

	defer db.Close()

	r := router.Router(db)
	r.Run("localhost:8080")
}
