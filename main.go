package Pokedex

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

var exit = make(chan struct{})

var cliCommandMap map[string]cliCommand

func initCommands() {
	cliCommandMap = map[string]cliCommand{
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
	exit <- struct{}{}
	return nil
}

func receiveCli(exit *chan struct{}) {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Pokedex > ")
			command, _ := reader.ReadString('\n')
			command = strings.TrimSpace(command)
			if cmd, ok := cliCommandMap[command]; ok {
				cmd.callback()
			}
		}
	}()
}

func main() {
	initCommands()
	receiveCli(&exit)
	<-exit
}
