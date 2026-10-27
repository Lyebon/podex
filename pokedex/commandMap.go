package main

import "fmt"

func commandMap(cfg *config) error{
	if cfg.next == nil && cfg.previous != nil{
		fmt.Println("This is the end of the list")
		return nil
	}

	areas, err := getResponse(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = areas.Next
	cfg.previous = areas.Previous
	for _, area:=range areas.Name{
		fmt.Println(area["name"])
	}
	return nil
}