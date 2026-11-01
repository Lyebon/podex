package main

import (
	"github.com/Lyebon/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	caught := make(map[string]pokeapi.PokeCatch)
	cfg := config{
		commands: getCommands(),
		client:   pokeClient,
		caught:	  caught,
	}
	startRepl(&cfg)
}
