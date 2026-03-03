package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DegsRed72/go_pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	caughtPokemon map[string]pokeapi.PokemonData
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := scanner.Text()
		cleanedWords := cleanInput(words)
		if len(cleanedWords) == 0 {
			continue
		}
		commandName := cleanedWords[0]
		args := cleanedWords[1:]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Cycles through 20 locations at a time",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Cycles backwards through 20 locations at a time",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Lists Pokemon that can be found in a given location (e.g. 'explore canalave-city-area')",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon of the given name (e.g. 'catch Pikachu')",
			callback:    commandCatch,
		},
	}
}
