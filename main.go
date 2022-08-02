package main

import (
	"my-app/src/routes"
)

func main() {
	router := routes.GetRouter()
	router.Run()
}
