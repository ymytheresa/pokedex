package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ymytheresa/pokedex/types"
)

type pokedexLocationAPI struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func PokedexGetLocation(cf *types.Config, next bool) {
	var url string
	if next {
		url = cf.NextUrl
	} else {
		url = cf.PrevUrl
	}

	if url == "" {
		fmt.Println("Reached first page")
		return
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var apiResults pokedexLocationAPI
	err = json.Unmarshal(body, &apiResults)
	if err != nil {
		log.Fatal(err)
	}

	cf.NextUrl = apiResults.Next
	cf.PrevUrl = apiResults.Previous
	for _, res := range apiResults.Results {
		fmt.Println(res.Name)
	}

}
