package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ymytheresa/pokedex/internal/pokeapi"
	"github.com/ymytheresa/pokedex/internal/pokecache"
	"github.com/ymytheresa/pokedex/types"
)

func commandCatch(cf *types.Config) error {
	name := cf.CmdSlice[1]
	url, baseExp, inCache := pokeapi.PokedexCatchPokemon(name)
	baseExpStringSlice := []string{strconv.Itoa(baseExp)}
	if !inCache {
		pokecache.GetPokedexCacheInstance().Add(url, baseExpStringSlice)
	}
	catchPokemonAndPrint(baseExp, name)
	return nil
}

func catchPokemonAndPrint(baseExp int, name string) {
	maxBase := 608 * 1.1 //Blissey's base experience 608, hightest amoung all pokemons, ensure 10% for this as a baseline
	rate := 1 - (float64(baseExp) / maxBase)

	var successRate float64
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Float64() <= rate {
		successRate = r.Float64()*50 + 50
	} else {
		successRate = r.Float64() * 50
	}

	printSlice := []string{fmt.Sprintf("Throwing a Pokeball at %s...", name)}
	if successRate > 50 {
		printSlice = append(printSlice, fmt.Sprintf("%s was caught!", name))
	} else {
		printSlice = append(printSlice, fmt.Sprintf("%s escaped!", name))
	}
	for _, name := range printSlice {
		fmt.Println(name)
	}
}
