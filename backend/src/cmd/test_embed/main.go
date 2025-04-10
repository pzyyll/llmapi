package main

import (
	"embed"
)

//go:embed *
var files embed.FS

func main() {
	// This is just a placeholder for the main function.
	// The actual functionality would be implemented in the backend code.
	data, _ := files.ReadFile("main.go")

	println(string(data))
}