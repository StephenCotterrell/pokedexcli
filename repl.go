package main

import "strings"

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(strings.ToLower(text))
	return cleanedInput
}
