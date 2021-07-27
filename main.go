package main

import (
	"github.com/joho/godotenv"
	"os"
	_ "platform-sample/docs"
	"platform-sample/infrastructure/database"
	"platform-sample/infrastructure/server"
)

func init() {
	godotenv.Load(".env.profile")
	godotenv.Load(".env." + os.Getenv("GO_PROFILE"))
}

// @title Platform-sample Swagger API
// @version 1.0.0
// @host localhost:8395
// @BasePath /api
func main() {
	db := database.SqlStore{}.GetDb()
	defer db.Close()
	server.Server{MainDb: db}.Init()
}
