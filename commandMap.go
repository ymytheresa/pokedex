package main

import (
	"github.com/ymytheresa/pokedex/internal/pokeapi"
	"github.com/ymytheresa/pokedex/types"
)

func commandMap(cf *types.Config) error {
	pokeapi.PokedexGetLocation(cf, true)
	return nil
}
