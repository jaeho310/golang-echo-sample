package main

import (
	"github.com/joho/godotenv"
	"os"
	"platform-sample/infrastructure/database"
	"platform-sample/infrastructure/server"
)


func init() {
	godotenv.Load(".env.profile")
	godotenv.Load(".env." + os.Getenv("GO_PROFILE"))
}

func main() {
	db := database.SqlStore{}.Getdb()
	defer db.Close()

	server.Server{MainDb: db}.Init()
}
