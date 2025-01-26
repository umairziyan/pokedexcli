package pokeapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
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

func (c *Client) GetLocations(pageUrl *string) (Locations, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}
	}
	req, err := http.NewRequest("GET", url, nil)
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

	return locations, nil
}
