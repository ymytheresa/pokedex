package main

import (
	"github.com/ymytheresa/pokedex/internal/pokeapi"
	"github.com/ymytheresa/pokedex/types"
)

func commandCatch(cf *types.Config) error {
	name := cf.CmdSlice[1]
	pokeapi.PokedexCatchPokemon(name)
	return nil
}
