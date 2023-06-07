package main

import (
	"github.com/sirupsen/logrus"
	"github.com/words/logger"
	"github.com/words/server"
)

func main() {
	logger.LogSetup()

	logrus.Printf("logrus main")

	server.Start()

}
