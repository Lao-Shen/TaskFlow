package main

import (
	"backend/internal"
	"fmt"
)

func main() {
	r := internal.ServerInit()

	err := r.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello, TaskFlow")
}
