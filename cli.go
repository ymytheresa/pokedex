package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ymytheresa/pokedex/internal/pokecache"
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
		"explore": {
			name:        "explore",
			description: "explore <location-area> to see list of pokemon of that location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch <pokemon name>, pokemon might be caught or escaped",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show all the caught pokemon",
			callback:    commandPokedex,
		},
	}
}

func receiveCli() {
	reader := bufio.NewReader(os.Stdin)
	cf := initConfig()
	pokecache.NewCache(5)
	for {
		fmt.Print("Pokedex > ")
		command, _ := reader.ReadString('\n')
		commandSlice := cleanCmd(command)
		if cmd, ok := getCliCommandMap()[commandSlice[0]]; ok {
			cf.CmdSlice = commandSlice
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
