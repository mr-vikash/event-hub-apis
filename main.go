package main

import (
	"eventhub/config"
	"eventhub/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	routes.RegisterRoutes()

	log.Println("Server is running on port 8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
