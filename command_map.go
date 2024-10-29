package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func mapCallback() error {
	url := "https://pokeapi.co/api/v2/location-area/"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error getting request: %w", err)
		return err
	}
	defer res.Body.Close()

	var LA LocationArea
	json.Unmarshal(res, &LA)

	return nil
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
