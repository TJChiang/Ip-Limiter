package main

import (
	"IpLimiter/cmd/internal"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	server, err := internal.InitializeServer()
	if err != nil {
		panic(err)
	}

	server.Run()
}
