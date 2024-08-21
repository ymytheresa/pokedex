package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/ymytheresa/pokedex/internal/pokeapi"
	"github.com/ymytheresa/pokedex/internal/pokecache"
	"github.com/ymytheresa/pokedex/types"
)

func commandCatch(cf *types.Config) error {
	name := cf.CmdSlice[1]
	url, baseExp, inCache := pokeapi.PokedexCatchPokemon(name)
	successChanceInStringSlice, successChanceInFloat := calcCatchSuccessRate(baseExp)
	if !inCache {
		pokecache.GetPokedexCacheInstance().Add(url, successChanceInStringSlice)
	}
	catchPokemonAndPrint(successChanceInFloat, name)
	return nil
}

func calcCatchSuccessRate(sBase int) ([]string, float64) {
	maxBase := 608 * 1.1 //Blissey's base experience 608, hightest amoung all pokemons, ensure 10% for this as a baseline
	fRate := 1 - (float64(sBase) / maxBase)
	fmt.Println(sBase, fRate)
	return []string{fmt.Sprintf("%f", fRate)}, fRate
}

func catchPokemonAndPrint(rate float64, name string) {
	var successRate float64
	if rand.Float64() <= rate/100 {
		successRate = rand.Float64()*50 + 50
	} else {
		successRate = rand.Float64() * 50
	}

	printSlice := []string{fmt.Sprintf("Throwing a Pokeball at %s...", name)}
	if successRate > 0.5 {
		printSlice = append(printSlice, fmt.Sprintf("%s was caught!", name))
	} else {
		printSlice = append(printSlice, fmt.Sprintf("%s escaped!", name))
	}
	for _, name := range printSlice {
		fmt.Println(name)
	}
}
