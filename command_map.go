package main

import (
	"fmt"
)

func commandMapf(cfg *config) error {
	locationsRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	for _, location := range locationsRes.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Printf("you're on the first page\n")
		return nil
	}

	locationsRes, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	for _, location := range locationsRes.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
