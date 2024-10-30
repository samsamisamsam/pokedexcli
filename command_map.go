package main

import (
	"fmt"
	"log"
)

func mapCallback(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("***LOCATIONS***")
	for _, area := range resp.Results {
		fmt.Printf("- %v\n", area.Name)
	}

	return nil
}
