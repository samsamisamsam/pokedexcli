package main

import (
	"errors"
	"fmt"
)

func mapCallback(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("***LOCATIONS***")
	for _, area := range resp.Results {
		fmt.Printf("- %v\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func mapbCallback(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("error: already at the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("***LOCATIONS***")
	for _, area := range resp.Results {
		fmt.Printf("- %v\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}
