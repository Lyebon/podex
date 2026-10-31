package main

import (
	"fmt"
)

func commandExplore (cfg *config, rest ...string) error {
	pokeEnc, err := cfg.client.GetPokemonEncounter(rest[0])
	if err != nil{
		return err
	}
	list := pokeEnc.PokemonEncounters
	fmt.Println("Exploring", rest[0])
	fmt.Println("Found Pokemon:")
	for _, value := range list{
		fmt.Println(" -", value.Pokemon.Name)
	}
	return nil
}