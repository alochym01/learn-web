package main

import (
	"fmt"

	"github.com/alochym01/learn-web/basic-layout/router"
)

func main() {
	fmt.Println("Server is starting with port 8080")
	router := router.Router()
	router.Run("localhost:8080")
}
