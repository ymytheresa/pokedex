package main

import (
	"github.com/ymytheresa/pokedex/internal/pokeapi"
	"github.com/ymytheresa/pokedex/types"
)

func commandExplore(cf *types.Config) error {
	name := cf.CmdSlice[1]
	pokeapi.PokedexGetPokemon(name)
	return nil
}
