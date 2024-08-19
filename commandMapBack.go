package main

import (
	"github.com/ymytheresa/pokedex/internal"
	"github.com/ymytheresa/pokedex/types"
)

func commandMapBack(cf *types.Config) error {
	internal.PokedexGetLocation(cf, false)
	return nil
}
