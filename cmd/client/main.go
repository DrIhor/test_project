package main

import (
	"log"

	cl "github.com/DrIhor/test_project/internal/client"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cl.StartWork()
}
