package main

import (
	"github.com/Anglepi/CommandFTL/api"
)

func main() {
	game := api.CreateServer()
	game.Start()
}