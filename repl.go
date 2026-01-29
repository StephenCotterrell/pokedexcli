package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var supportedCommands map[string]cliCommand

func init() {
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
			if err := command.callback(); err != nil {
				fmt.Print("there was an error: %w", err)
			}
		}
	}
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex! \nUsage: \n\n")
	for command := range supportedCommands {
		fmt.Printf("%s: %s\n", supportedCommands[command].name, supportedCommands[command].description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(strings.ToLower(text))
	return cleanedInput
}
