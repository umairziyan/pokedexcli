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

func GetLocations(pageUrl *string) (Locations, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return Locations{}, err
	}

	var locations Locations
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		log.Fatal(err)
		return Locations{}, err
	}

	return locations, nil
}
