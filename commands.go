package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
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

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	fmt.Printf("Exploring %v...\n", args[0])
	fmt.Print("Found Pokemon:\n")

	alldetails, err := cfg.client.GetPokemonList(args[0])
	if err != nil {
		return err
	}

	for _, pokemon := range alldetails.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("you must provide only one pokemon")
	}

	pokemonDetails, err := cfg.client.GetPokemonDetails(args[0])
	if err != nil {
		return err
	}

	pokemonName := pokemonDetails.Name
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	// Determine a catch probablility
	baseExperience := pokemonDetails.BaseExperience
	maxBaseExperience := 255
	minCatchChance := 0.1
	maxCatchChance := 0.9
	catchProbability := maxCatchChance - ((float64(baseExperience) / float64(maxBaseExperience)) * (maxCatchChance - minCatchChance))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomValue := r.Float64()

	if randomValue <= catchProbability {
		fmt.Printf("%v was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonDetails.Name] = pokemonDetails
		return nil
	}
	fmt.Printf("%v escaped!\n", pokemonName)

	return nil
}
