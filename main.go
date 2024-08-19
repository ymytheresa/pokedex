package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	var output strings.Builder
	output.WriteString("Welcome to the Pokedex!\n")
	output.WriteString("Usage:\n\n")

	for name, info := range cliCommandMap {
		fmt.Fprintf(&output, "%-10s: %s\n", name, info.description)
	}
	output.WriteString("\n")

	// Print the final result
	fmt.Println(output.String())

	return nil
}

func commandExit() error {
	// Implementation here
	return nil
}

var cliCommandMap = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func receiveCli(exit *ch) error {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Pokedex >")
			command, _ := reader.ReadString('\n')
			command = strings.TrimSpace(command)

		}
	}()
}

func main() {
	exit := make(ch, struct{})
	exit <- receiveCli(exit)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s!\n", name)
}
