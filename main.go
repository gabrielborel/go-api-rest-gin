package main

import (
	"api-rest-gin/database"
	"api-rest-gin/routes"
)

func main() {
	database.ConnectWithDatabase()
	
	routes.HandleRoutes()
}
