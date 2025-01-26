package main

import (
	"time"

	"github.com/umairziyan/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{client: pokeClient}
	replStart(cfg)
}
