package main

import (
	"fmt"

	"github.com/StephenCotterrell/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, arguments []string) error {
	pokemonName := arguments[0]

	if pokemonData, seenBefore := cfg.pokedex[pokemonName]; seenBefore {
		stats := statsToMap(pokemonData.Stats)
		types := typesToSlice(pokemonData.Types)

		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n  -hp: %d\n  -attack: %d\n  -defense: %d\n  -special-attack: %d\n  -special-defense: %d\n  -speed: %d\n", pokemonName, pokemonData.Height, pokemonData.Weight, stats["hp"], stats["attack"], stats["defense"], stats["special-attack"], stats["special-defense"], stats["speed"])
		fmt.Printf("Types:\n")
		for _, t := range types {
			fmt.Printf("  - %s\n", t)
		}
	} else {
		fmt.Printf("you have not caught that pokemon")
	}

	return nil
}

func statsToMap(stats []pokeapi.PokemonStats) map[string]int {
	m := make(map[string]int, len(stats))
	for _, s := range stats {
		m[s.Stat.Name] = int(s.BaseStat)
	}
	return m
}

func typesToSlice(types []pokeapi.PokemonTypes) []string {
	t := []string{}
	for _, s := range types {
		t = append(t, s.Type.Name)
	}

	return t
}
