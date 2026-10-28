package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"github.com/Lyebon/pokedex/internal/pokecache"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for true{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0{
			continue
		}
		commandName := words[0]

		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct{
	commands map[string]cliCommand
	pokeApi		string
	next		*string
	previous	*string
	cache		*pokecache.Cache
}

type pokemonAreas struct{
	Next		*string `json:"next"`
	Previous	*string `json:"previous"`
	Name		[]map[string]string `json:"results"`
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
	    "exit": {
    	    name:        "exit",
        	description: "Exit the Pokedex",
        	callback:    commandExit,
    	},
		"help": {
			name:		 "help",
			description: "Display a help message",
			callback:	 commandHelp,
		},
		"map":{
			name:		 "map",
			description: "Display the next 20 locations",
			callback:	 commandMap,
		},
		"mapb":{
			name:		 "mapb",
			description: "Display the previous 20 locations",
			callback:	 commandMapB,
		},
	}
}
