package main

import (
	"os"

	"github.com/ymytheresa/pokedex/types"
)

func commandExit(_ *types.Config) error {
	os.Exit(0)
	return nil
}
