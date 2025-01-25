package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		output := cleanInput(scanner.Text())
		if len(output) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", output[0])
	}
}

func cleanInput(text string) []string {
	// A function to split strings by white space and return them in a slice in lowercase.
	v := strings.Fields(strings.ToLower(text))

	return v
}
