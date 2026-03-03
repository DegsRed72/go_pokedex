package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Must include location name when entering 'explore' command")
	}
	fmt.Println("Exploring...")

	locationData, err := cfg.pokeapiClient.ListData(&args[0])
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")

	for _, pokemon := range locationData.PokemonEncounters {
		fmt.Printf("- %s", pokemon.Pokemon.Name)
		fmt.Println()
	}

	return nil
}
