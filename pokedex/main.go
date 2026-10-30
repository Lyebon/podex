package main

import (
	"github.com/Lyebon/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := config{
		commands: getCommands(),
		client:   pokeClient,
	}
	startRepl(&cfg)
}
