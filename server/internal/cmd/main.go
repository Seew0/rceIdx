package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/seew0/rceIdx/internal/cmd/server"
)

var (
	version int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error occured while loading .env file: %v", err)
	}
}

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()

	Server := server.NewServer(port, router)
	Server.Run()
}
