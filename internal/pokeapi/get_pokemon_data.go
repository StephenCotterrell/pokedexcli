package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(name string) (PokemonData, error) {
	url := baseURL + "/pokemon/" + name

	if data, exists := c.cache.Get(url); exists {
		pokemonDataRes := PokemonData{}
		err := json.Unmarshal(data, &pokemonDataRes)
		if err != nil {
			return PokemonData{}, err
		}

		return pokemonDataRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}

	if res.StatusCode != http.StatusOK {
		return PokemonData{}, fmt.Errorf("response had non-ok status code: %d", res.StatusCode)
	}

	defer func() {
		if cerr := res.Body.Close(); err != nil && cerr != nil {
			err = cerr
		}
	}()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{}, err
	}

	c.cache.Add(url, data)

	pokemonDataRes := PokemonData{}

	err = json.Unmarshal(data, &pokemonDataRes)
	if err != nil {
		return PokemonData{}, err
	}

	return pokemonDataRes, nil
}
