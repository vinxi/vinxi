package main

import (
	"fmt"
	"os"

	"gopkg.in/vinxi/vinxi.v0"
	"gopkg.in/vinxi/vinxi.v0/sandbox"
	"gopkg.in/vinxi/vinxi.v0/sandbox/plugins/static"
	"gopkg.in/vinxi/vinxi.v0/sandbox/rules"
)

const port = 3100

func main() {
	cwd, _ := os.Getwd()

	// Create a new vinxi proxy
	v := vinxi.New()

	// Manage current vinxi instance
	manager := sandbox.Manage(v)
	scope := manager.NewScope(rules.NewPath("/foo"))
	scope.UsePlugin(static.New(cwd))

	go func() {
		fmt.Printf("Admin server listening on port: %d\n", port+100)
		manager.ServeAndListen(sandbox.ServerOptions{Port: port + 100})
	}()

	// Target server to forward
	v.Forward("http://httpbin.org")

	fmt.Printf("Server listening on port: %d\n", port)
	_, err := v.ServeAndListen(vinxi.ServerOptions{Port: port})
	if err != nil {
		fmt.Errorf("Error: %s\n", err)
	}
}
