package main

import "fmt"

func commandHelp(config *config, arguments []string) error {
	fmt.Printf("Welcome to the Pokedex! \nUsage: \n\n")
	supportedCommands := getCommands()
	for command := range supportedCommands {
		fmt.Printf("%s: %s\n", supportedCommands[command].name, supportedCommands[command].description)
	}
	return nil
}
