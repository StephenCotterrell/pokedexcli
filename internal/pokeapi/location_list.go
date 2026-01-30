package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaList{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaList{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	locationsRes := LocationAreaList{}

	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return LocationAreaList{}, err
	}

	return locationsRes, nil
}
