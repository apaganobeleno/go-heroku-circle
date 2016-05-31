package main

import (
	"log"
	"net/http"
	"os"
	"templo/db"
	"templo/routing"

	"github.com/codegangsta/negroni"
	"github.com/joho/godotenv"
)

var port = portWithDefault("3000")

func main() {
	loadEnvironmentFromFile()
	checkDBConnection()

	router := routing.BuildRouter()
	n := negroni.Classic()
	n.UseHandler(router)

	if os.Getenv("DEVELOPMENT") == "1" {
		log.Println("| WARNING: App running in development mode")
	}

	log.Printf("| App running on port %v", port)
	http.ListenAndServe(":"+port, n)
}

func checkDBConnection() {
	_, err := db.Connection()
	if err != nil {
		log.Println("| WARNING: Could not connect to the DATABASE_URL set")
		log.Printf("| ERROR %v", err.Error())
	}
}

func loadEnvironmentFromFile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("| WARNING: Could not find .env file relaying on system ENV")
	}
}

func portWithDefault(alternative string) string {
	port := os.Getenv("PORT")
	if port == "" {
		return alternative
	}
	return port
}
