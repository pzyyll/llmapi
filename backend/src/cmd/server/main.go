package main

import (
	"llmapi/src/internal/core"
)

func main() {
	if err := core.InitServer(); err != nil {
		panic(err)
	}

	core.Run()
}
