package main

import "fmt"

func commandMapB(cfg *config) error{
	if cfg.previous == nil{
		fmt.Println("This is the start of the list")
		return nil
	}
	areas, err := getResponse(cfg.previous, cfg)
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