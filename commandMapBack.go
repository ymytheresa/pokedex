package main

import (
	"github.com/ymytheresa/pokedex/internal/pokeapi"
	"github.com/ymytheresa/pokedex/types"
)

func commandMapBack(cf *types.Config) error {
	pokeapi.PokedexGetLocation(cf, false)
	return nil
}
