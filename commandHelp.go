package main

import (
	"fmt"
	"strings"
)

func commandHelp() error {
	var output strings.Builder
	output.WriteString("Welcome to the Pokedex!\n")
	output.WriteString("Usage:\n\n")

	for name, info := range getCliCommandMap() {
		fmt.Fprintf(&output, "%-10s: %s\n", name, info.description)
	}
	output.WriteString("\n")

	fmt.Println(output.String())

	return nil
}
