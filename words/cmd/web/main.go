package main

import (
	"github.com/sirupsen/logrus"
	"github.com/words/cmd/server"
	"github.com/words/pkg/logger"
)

func main() {
	logger.LogSetupConsole() //настройка логгера

	logrus.Printf("Started server localhost:8080/")
	server.Start()

}
