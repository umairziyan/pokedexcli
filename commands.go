package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, args []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	for _, value := range getCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	return nil
}

func commandMap(cfg *config, args []string) error {
	locations, err := cfg.client.GetLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locations.Previous
	cfg.NextLocationsURL = locations.Next

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	locations, err := cfg.client.GetLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locations.Previous
	cfg.NextLocationsURL = locations.Next

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
