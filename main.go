package main

import (
	"github.com/br-jeff/go-lang-api/config"
	"github.com/br-jeff/go-lang-api/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.ErrorF("Erro on config initialization %v", err)
	}
	router.Initialize()
}
