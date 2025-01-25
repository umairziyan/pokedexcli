package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Hello, World!")
}

func cleanInput(text string) []string {
	// A function to split strings by white space and return them in a slice in lowercase.
	v := strings.Fields(strings.ToLower(text))

	return v
}
