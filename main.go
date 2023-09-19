package main

import (
	"github.com/luissena/golang-crud/config"
	"github.com/luissena/golang-crud/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("Config Initialization -> %v", err)
		return
	}

	router.Initialize()

}
