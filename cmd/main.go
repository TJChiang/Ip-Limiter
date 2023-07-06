package main

import (
	"IpLimiter/cmd/internal"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	godotenv.Load()
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
}

func main() {
	server, err := internal.InitializeServer()
	if err != nil {
		panic(err)
	}

	server.Run()
}
