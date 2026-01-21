package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := scanner.Text()
		cleaned_words := cleanInput(words)
		first_word := ""
		if len(cleaned_words) != 0 {
			first_word = cleaned_words[0]
		}
		fmt.Printf("Your command was: %s\n", first_word)
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
