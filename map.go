package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	mapData, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = mapData.Next
	cfg.Previous = mapData.Previous

	for _, loc := range mapData.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page.")
		return nil
	}
	mapData, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = mapData.Next
	cfg.Previous = mapData.Previous
	for _, loc := range mapData.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
