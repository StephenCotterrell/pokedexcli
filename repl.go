package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/StephenCotterrell/pokedexcli/internal/pokeapi"
)

var (
	supportedCommands map[string]cliCommand
	cfg               Config
)

func init() {
	next := pokeapi.BaseLocationAreaURL
	cfg.Next = &next
	cfg.Previous = nil

	supportedCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandGetNextLocationAreas,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the next 20 location areas",
			callback:    commandGetPreviousLocationAreas,
		},
	}
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()

		command, ok := supportedCommands[input]

		if !ok {
			fmt.Printf("Unknown command\n")
		} else {
			if err := command.callback(&cfg); err != nil {
				fmt.Printf("there was an error: %w", err)
			}
		}
	}
}

func commandGetNextLocationAreas(config *Config) error {
	LocationAreas, err := pokeapi.GetLocationAreas(*config.Next)
	if err != nil {
		log.Fatal(err)
	}

	config.Next = LocationAreas.Next
	config.Previous = LocationAreas.Previous

	for _, location := range LocationAreas.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandGetPreviousLocationAreas(config *Config) error {
	if config.Previous == nil {
		fmt.Printf("you're on the first page\n")
		return nil
	}

	LocationAreas, err := pokeapi.GetLocationAreas(*config.Previous)
	if err != nil {
		log.Fatal(err)
	}

	config.Next = LocationAreas.Next
	config.Previous = LocationAreas.Previous

	for _, location := range LocationAreas.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandExit(config *Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Printf("Welcome to the Pokedex! \nUsage: \n\n")
	for command := range supportedCommands {
		fmt.Printf("%s: %s\n", supportedCommands[command].name, supportedCommands[command].description)
	}
	return nil
}

type Config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(strings.ToLower(text))
	return cleanedInput
}
