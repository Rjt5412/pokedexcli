package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rjt5412/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display the help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"mapf": {
			name:        "mapf",
			description: "Display the names of next 20 location areas in the pokemon world based on current location",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of previous 20 locations based on current location",
			callback:    commandMapb,
		},
	}
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex >")
		reader.Scan()

		input := reader.Text()
		command, ok := commands[input]

		if ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Invalid command")
			commands["help"].callback(cfg)
		}

	}
}
