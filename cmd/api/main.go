package main

import (
	"fmt"
	"go-blueprint-htmx/internal/server"
)

func main() {
	fmt.Println("Starting server on :8080")
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
