package main

import (
	"github.com/Anglepi/CommandFTL/api"
	"github.com/Anglepi/CommandFTL/logger"
)

func main() {
	logger.InitLog()
	game := api.CreateServer()
	game.Start()
	defer logger.CloseLog()
}
