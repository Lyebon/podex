package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config, rest ...string) error {
	target := rest[0]
	pCatch, err := cfg.client.GetPokemonCatch(target)
	if err != nil{
		return err
	}
	random := rand.Intn(pCatch.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %v...", target)
	cal := int(math.Round(float64(pCatch.BaseExperience)/1.5))
	if random >= cal {
		fmt.Println(target, "was caught!")
		cfg.caught[target] = pCatch
		return nil
	} 

	fmt.Println(target, "escaped!")
	return nil
}