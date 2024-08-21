package main

import (
	"github.com/ymytheresa/pokedex/internal/pokeapi"
	"github.com/ymytheresa/pokedex/types"
)

func commandPokedex(cf *types.Config) error {
	pokeapi.FetchCaughtList()
	return nil
}
