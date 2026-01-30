package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if data, exists := c.cache.Get(url); exists {
		locationsRes := LocationAreaList{}
		err := json.Unmarshal(data, &locationsRes)
		if err != nil {
			return LocationAreaList{}, err
		}

		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaList{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaList{}, err
	}

	if res.StatusCode != http.StatusOK {
		return LocationAreaList{}, fmt.Errorf("response had non-ok status code: %d", res.StatusCode)
	}
	// defer res.Body.Close()

	defer func() {
		if cerr := res.Body.Close(); err != nil && cerr != nil {
			err = cerr
		}
	}()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	c.cache.Add(url, data)

	locationsRes := LocationAreaList{}

	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return LocationAreaList{}, err
	}

	return locationsRes, nil
}
