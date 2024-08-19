package main

import (
	"github.com/ymytheresa/pokedex/internal"
	"github.com/ymytheresa/pokedex/types"
)

func commandMap(cf *types.Config) error {
	internal.PokedexGetLocation(cf, true)
	return nil
}
