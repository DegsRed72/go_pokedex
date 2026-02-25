package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Map struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	res, err := http.Get(url)

	if err != nil {
		return errors.New("Error occurred while getting map.")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Error occurred with status code: %v\n", res.StatusCode)
	}
	if err != nil {
		return errors.New("Error occurred while getting map.")
	}

	mapData := Map{}
	erro := json.Unmarshal(body, &mapData)
	if erro != nil {
		return errors.New("Error occuured while unmarshalling map.")
	}
	for _, loc := range mapData.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = mapData.Next
	cfg.Previous = mapData.Previous

	return nil
}

func commandMapB(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Previous != nil {
		url = *cfg.Previous
	} else {
		fmt.Println("You're on the first page.")
		return nil
	}
	res, err := http.Get(url)

	if err != nil {
		return errors.New("Error occurred while getting map.")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Error occurred with status code: %v\n", res.StatusCode)
	}
	if err != nil {
		return errors.New("Error occurred while getting map.")
	}

	mapData := Map{}
	erro := json.Unmarshal(body, &mapData)
	if erro != nil {
		return errors.New("Error occuured while unmarshalling map.")
	}
	for _, loc := range mapData.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = mapData.Next
	cfg.Previous = mapData.Previous

	return nil
}
