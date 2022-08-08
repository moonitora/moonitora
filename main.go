package main

import (
	"fmt"
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/router"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if err := database.Connect(); err != nil {
		log.Fatalf("Failed while connecting to database: %v", err.Error())
		return
	}

	engine := router.Build()
	engine.Run(fmt.Sprintf(":%s", port))
}
