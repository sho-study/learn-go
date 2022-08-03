package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"my-app/src/database"
	"my-app/src/routes"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("環境変数の読み込みに失敗")
	}
}

func init() {
	loadEnv()
	database.Connect()
	database.Migrate()
}

func main() {
	router := routes.GetRouter()
	router.Run()
}
