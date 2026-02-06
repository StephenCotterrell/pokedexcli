package main

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/StephenCotterrell/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, arguments []string) error {
	pokemonName := arguments[0]

	fmt.Printf("\nThrowing a Pokeball at %s...\n", pokemonName)
	if _, alreadyCaught := cfg.pokedex[pokemonName]; alreadyCaught {
		fmt.Printf("You have already caught this pokemon!\n")
		return nil
	}

	pokemonDataRes, err := cfg.pokeapiClient.GetPokemonData(pokemonName)
	if err != nil {
		return err
	}

	pokemonBaseExperience := pokemonDataRes.BaseExperience

	caught := rand.Float64() < catchProb(int(pokemonBaseExperience))

	if caught {
		fmt.Printf("%s was caught!\n", pokemonName)
		if cfg.pokedex == nil {
			cfg.pokedex = make(map[string]pokeapi.PokemonData)
		}
		cfg.pokedex[pokemonName] = pokemonDataRes
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func catchProb(baseExp int) float64 {
	b := float64(baseExp)

	pMin := 0.20
	pMax := 0.60
	b0 := 112.0
	k := 35.0

	s := 1.0 / (1.0 + math.Exp((b-b0)/k))
	return pMin + (pMax-pMin)*s
}
