package main

import "fmt"

func commandPokedex(cfg *config, arguments []string) error {
	if len(cfg.pokedex) > 0 {
		fmt.Printf("Your Pokedex:\n")
		for _, pokemon := range cfg.pokedex {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	} else {
		fmt.Printf("you have not caught any pokemon yet")
	}

	return nil
}
