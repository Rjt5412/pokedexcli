package main

import (
	"time"

	"github.com/rjt5412/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Hour / 2)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
