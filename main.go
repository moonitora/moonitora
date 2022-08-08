package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if err := database.Connect(); err != nil {
		log.Fatalf("Failed while connecting to database: %v", err.Error())
		return
	}
}
