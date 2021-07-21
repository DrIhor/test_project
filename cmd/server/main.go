package main

import (
	"fmt"
	"log"

	sr "github.com/DrIhor/test_project/internal/service/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("Start server")
	sr.StartServer()
}
