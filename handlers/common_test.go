package handlers

import (
	"log"

	"github.com/CloudyKit/jet"
	"github.com/joho/godotenv"
)

func init() {
	//For testing proposes
	views = jet.NewHTMLSet("../views")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("| WARNING: Could not find .env file relaying on system ENV")
	}
}
