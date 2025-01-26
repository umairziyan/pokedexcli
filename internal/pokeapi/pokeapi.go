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

func (c *Client) GetPokemonList(location string) (LocationDetails, error) {
	URL := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(URL); ok {
		locationsResp := LocationDetails{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationDetails{}, err
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
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	var locationDetails LocationDetails
	err = json.Unmarshal(dat, &locationDetails)
	if err != nil {
		log.Fatal(err)
		return LocationDetails{}, err
	}

	c.cache.Add(URL, dat)
	return locationDetails, nil
}

func (c *Client) GetPokemonDetails(pokemonName string) (Pokemon, error) {
	URL := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(URL); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalf("error request function: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalf("error response function: %v", err)
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		log.Fatal(err)
		return Pokemon{}, err
	}

	c.cache.Add(URL, dat)
	return pokemon, nil
}
