package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	URL := baseURL + "/location-area"
	if pageURL != nil {
		URL = *pageURL
	}

	if val, ok := c.cache.Get(URL); ok {
		locationsResp := Locations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Locations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalf("error request function: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalf("error response function: %v", err)
		return Locations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	var locations Locations
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		log.Fatal(err)
		return Locations{}, err
	}

	c.cache.Add(URL, dat)
	return locations, nil
}
