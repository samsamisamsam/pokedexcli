package main

import (
	"fmt"
	"net/http"
)

func mapCallback() error {
	url := "https://pokeapi.co/api/v2/location/15/"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error getting request: %w", err)
		return err
	}
	defer res.Body.Close()

	fmt.Println(res)

	return nil
}

/*
type location struct {
	id   int
	name string
}

/*
Location (type)
Name	Description	Type

id
The identifier for this resource.
integer

name
The name for this resource.
string

region
The region this location can be found in.
NamedAPIResource (Region)

names
The name of this resource listed in different languages.
list Name

game_indices
A list of game indices relevent to this location by generation.
list GenerationGameIndex

areas
Areas that can be found within this location.
list NamedAPIResource (LocationArea) */
