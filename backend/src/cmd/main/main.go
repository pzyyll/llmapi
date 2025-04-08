package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file", err)
		return
	}
	fmt.Println("Hello, World!")
}
