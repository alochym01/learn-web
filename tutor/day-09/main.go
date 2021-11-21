package main

import "github.com/alochym01/web-w-gin/router"

func main() {
	r := router.Router()
	r.Run("localhost:8080")
}
