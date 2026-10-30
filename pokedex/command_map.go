package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	if cfg.nextLocationsURL == nil && cfg.prevLocationsURL != nil {
		fmt.Println("This is the end of the list")
		return nil
	}

	areas, err := cfg.client.GetResponse(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = areas.Next
	cfg.prevLocationsURL = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("This is the start of the list")
		return nil
	}
	areas, err := cfg.client.GetResponse(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = areas.Next
	cfg.prevLocationsURL = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}
