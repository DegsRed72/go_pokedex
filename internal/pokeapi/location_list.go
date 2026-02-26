package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (Map, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	dat, ok := c.cache.Get(url)
	if !ok {
		var err error
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Map{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Map{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return Map{}, err
		}
		c.cache.Add(url, dat)
	}

	locationsResp := Map{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Map{}, err
	}

	return locationsResp, nil
}
