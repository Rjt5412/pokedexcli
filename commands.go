package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandMapf(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.previousLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationURL == nil {
		return errors.New("you are on the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.previousLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(c *config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]
	location, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf("%v \n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	const threshold = 50
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s \n...", pokemonName)
	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	ranNum := rand.Intn(pokemon.BaseExperience)

	if ranNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught! \n You can now inspect it with the inspect command. \n", pokemonName)
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s \n", pokemon.Name)
	fmt.Printf("Height: %v \n", pokemon.Height)
	fmt.Printf("Weight: %v \n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s:%d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  -%s\n", typeInfo.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {

	caughtPokemon := cfg.caughtPokemon

	if len(caughtPokemon) == 0 {
		return fmt.Errorf("pokedex is empty. You haven't caught any pokemon so far")
	}

	fmt.Println("Your Pokedex:")
	for name := range caughtPokemon {
		fmt.Println("  -", name)
	}
	return nil
}
