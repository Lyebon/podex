package main

import (
	"time"
    "github.com/Lyebon/pokedex/internal/pokecache"
)
func main(){
	cfg:= config{
		commands: 	getCommands(),
		pokeApi:	"https://pokeapi.co/api/v2/location-area/",
		cache:		pokecache.NewCache(5 * time.Minute),
	}
	startRepl(&cfg)
}