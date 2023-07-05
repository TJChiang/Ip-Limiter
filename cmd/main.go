package main

import "IpLimiter/cmd/internal"

func main() {
	server, err := internal.InitializeServer()
	if err != nil {
		panic(err)
	}

	server.Run()
}
