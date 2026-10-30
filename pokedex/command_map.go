package main

import (
	"fmt"
)

func commandMap(cfg *config, rest ...string) error {
	areas, err := cfg.client.GetAreaResponse(cfg.nextLocationsURL)
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

func commandMapB(cfg *config, rest ...string) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("This is the start of the list")
		return nil
	}
	areas, err := cfg.client.GetAreaResponse(cfg.prevLocationsURL)
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
