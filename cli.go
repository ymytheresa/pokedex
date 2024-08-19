package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ymytheresa/pokedex/types"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*types.Config) error
}

func initConfig() *types.Config {
	return &types.Config{
		NextUrl: "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0",
	}
}

func getCliCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Show page of map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show last page of map",
			callback:    commandMapBack,
		},
	}
}

func receiveCli() {
	reader := bufio.NewReader(os.Stdin)
	cf := initConfig()
	for {
		fmt.Print("Pokedex > ")
		command, _ := reader.ReadString('\n')
		commandSlice := cleanCmd(command)
		if cmd, ok := getCliCommandMap()[commandSlice[0]]; ok {
			err := cmd.callback(cf)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanCmd(cmd string) []string {
	//return slice of lower case string that the substring were separated by whitespave as delimiter
	return strings.Fields(strings.ToLower(strings.TrimSpace(cmd)))
}