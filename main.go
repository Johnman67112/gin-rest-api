package main

import (
	"github.com/Johnman67112/gin-rest-api/database"
	"github.com/Johnman67112/gin-rest-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	database.DatabaseConnect()
	routes.HandleRequests()
}
