package pokeapi

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

var httpCli = NewClient(10)

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

	request, err := http.NewRequest(url)
	if err != nil {
		log.Fatal(err)
	}

	response, err := httpCli.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
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
