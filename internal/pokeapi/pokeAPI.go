package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ymytheresa/pokedex/internal/pokecache"
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

var httpCli Client = NewClient(100 * time.Second)

func PokedexGetLocation(cf *types.Config, next bool) {
	var url string

	if next {
		url = cf.NextUrl
	} else {
		url = cf.PrevUrl
	}

	exist := checkAndPrintFromMem(url)
	if exist {
		return
	}

	addToMemAndPrint(url, cf)
}

func addToMemAndPrint(url string, cf *types.Config) {
	if url == "" {
		fmt.Println("Reached first page")
		return
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := httpCli.httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, body)
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
	cacheEntryVal := make([]string, 20)
	for _, res := range apiResults.Results {
		cacheEntryVal = append(cacheEntryVal, res.Name)
	}
	pokecache.GetPokedexCacheInstance().Add(url, cacheEntryVal)
	printLocation(cacheEntryVal)
}

func checkAndPrintFromMem(url string) bool {
	val, ok := pokecache.GetPokedexCacheInstance().Get(url)
	if !ok {
		return false
	}
	printLocation(val)
	return true
}

func printLocation(val []string) {
	for _, name := range val {
		fmt.Println(name)
	}
}
