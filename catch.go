package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	pokemon := &args[0]
	fmt.Printf("Throwing a Pokeball at %s...", *pokemon)
	fmt.Println()
	pokemonData, err := cfg.pokeapiClient.ListPokemonData(pokemon)
	if err != nil {
		return err
	}
	baseXP := pokemonData.BaseExperience
	catchDifficulty := math.Pow(float64(baseXP), 0.7)
	if rand.Intn(100) > int(catchDifficulty) {
		fmt.Printf("%s was caught!", *pokemon)
		cfg.caughtPokemon[*pokemon] = pokemonData
		fmt.Println()
		fmt.Println("You may now inspect this Pokemon with the inspect command.")
	} else {
		fmt.Printf("%s escaped!", *pokemon)
	}
	fmt.Println()

	return nil
}
