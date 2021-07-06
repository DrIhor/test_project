package main

import (
	"fmt"

	sr "github.com/DrIhor/test_project/internal/server"
)

func main() {
	fmt.Println("Start server")
	sr.StartServer()
}
