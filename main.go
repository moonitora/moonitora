package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/router"
	"github.com/victorbetoni/moonitora/security"
	"log"
	"os"
)

func init() {
	str, _ := security.Hash("p%cUa7v6A3e&5x9cnO0v2&G7")
	fmt.Println(string(str))
}

func main() {
	port := os.Getenv("PORT")

	if err := database.Connect(); err != nil {
		log.Fatalf("Failed while connecting to database: %v", err.Error())
		return
	}

	engine := router.Setup(gin.Default())
	engine.Run(fmt.Sprintf(":%s", port))
}
