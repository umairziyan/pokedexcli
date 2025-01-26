package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/umairziyan/pokedexcli/internal/pokeapi"
)

func replStart(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		args := input[1:]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

type config struct {
	client           pokeapi.Client
	NextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		}, "map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get a list of pokemon from a specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch <arg> - gives you a chance at catching a pokemon",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	// A function to split strings by white space and return them in a slice in lowercase.
	v := strings.Fields(strings.ToLower(text))

	return v
}
