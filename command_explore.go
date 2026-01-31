package main

import "fmt"

func commandExplore(cfg *config, arguments []string) error {
	location := arguments[0]
	fmt.Printf("Exploring %s...\n", location)
	fmt.Printf("Found Pokemon: \n")

	pokemonEncountersRes, err := cfg.pokeapiClient.ExploreLocationPokemon(location)
	if err != nil {
		return err
	}

	for _, pokemonEncounter := range pokemonEncountersRes.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
