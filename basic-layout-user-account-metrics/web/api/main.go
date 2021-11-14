package main

import (
	"database/sql"
	"expvar"
	"fmt"
	"runtime"

	"github.com/alochym01/learn-web/basic-layout-user-account/router"
)

func main() {
	fmt.Println("Server is starting with port 8080")
	db, err1 := sql.Open("sqlite3", "alochym.db")

	if err1 != nil {
		panic(err1)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	expvar.Publish("db", expvar.Func(func() interface{} {
		return db.Stats()
	}))

	defer db.Close()

	expvar.Publish("goroutines", expvar.Func(func() interface{} {
		return runtime.NumGoroutine()
	}))
	// logger := logger.NewLogger()
	router := router.Router(db)
	router.Run("localhost:8080")
}
