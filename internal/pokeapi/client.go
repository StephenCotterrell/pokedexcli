package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaList struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

const BaseLocationAreaURL string = "https://pokeapi.co/api/v2/location-area"

func GetLocationAreas(url string) (LocationAreaList, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaList{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return LocationAreaList{}, fmt.Errorf("pokeapi returned status %d: %s; body: %s", res.StatusCode, res.Status, string(body))
	}

	locationAreas := LocationAreaList{}

	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return LocationAreaList{}, err
	}

	return locationAreas, nil
}
