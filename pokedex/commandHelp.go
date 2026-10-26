package main

import (
	"fmt"
)

func commandHelp(cfg *config) error{
	fmt.Printf(`
	Welcome to the Pokedex!
Usage:
`)
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}