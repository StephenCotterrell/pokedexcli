package main

import "strings"

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(text)
	return cleanedInput
}
