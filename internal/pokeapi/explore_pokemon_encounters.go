package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocationPokemon(location string) (PokemonEncountersList, error) {
	url := baseURL + "/location-area/" + location

	if data, exists := c.cache.Get(url); exists {
		pokemonEncountersRes := PokemonEncountersList{}
		err := json.Unmarshal(data, &pokemonEncountersRes)
		if err != nil {
			return PokemonEncountersList{}, err
		}

		return pokemonEncountersRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncountersList{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonEncountersList{}, err
	}

	if res.StatusCode != http.StatusOK {
		return PokemonEncountersList{}, fmt.Errorf("response had non-ok status code: %d", res.StatusCode)
	}
	// defer res.Body.Close()

	defer func() {
		if cerr := res.Body.Close(); err != nil && cerr != nil {
			err = cerr
		}
	}()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonEncountersList{}, err
	}

	c.cache.Add(url, data)

	pokemonEncountersRes := PokemonEncountersList{}

	err = json.Unmarshal(data, &pokemonEncountersRes)
	if err != nil {
		return PokemonEncountersList{}, err
	}

	return pokemonEncountersRes, nil
}
