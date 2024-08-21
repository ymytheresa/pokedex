package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	"github.com/ymytheresa/pokedex/internal/pokecache"
	"github.com/ymytheresa/pokedex/types"
)

var httpCli Client = NewClient(100 * time.Second)

func PokedexGetLocation(cf *types.Config, next bool) {
	var url string

	if next {
		url = cf.NextUrl
	} else {
		url = cf.PrevUrl
	}

	_, ok := checkAndPrintFromMem(url, true)
	if ok {
		return
	}

	addToMemAndPrintLocation(url, cf)
}

func callAPI(url string) ([]byte, error) {
	if url == "" {
		return []byte{}, fmt.Errorf("reached first page")
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	response, err := httpCli.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if response.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and body: %s", response.StatusCode, body)
	}

	return body, nil
}

func addToMemAndPrintLocation(url string, cf *types.Config) {
	body, err := callAPI(url)
	if err != nil {
		log.Fatal(err)
	}

	var apiResults types.PokedexLocationAPI
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
	printStringSlice(cacheEntryVal)
}

func checkAndPrintFromMem(url string, print bool) ([]string, bool) {
	val, ok := pokecache.GetPokedexCacheInstance().Get(url)
	if !ok {
		return []string{}, false
	}
	if print {
		printStringSlice(val)
	}
	return val, true
}

func printStringSlice(val []string) {
	for _, name := range val {
		fmt.Println(name)
	}
}

func PokedexGetPokemon(name string) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)

	_, ok := checkAndPrintFromMem(url, true)
	if ok {
		return
	}

	body, err := callAPI(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	var apiResults types.PokedexPokemonByLocationAPI
	err = json.Unmarshal(body, &apiResults)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	cacheEntryVal := make([]string, 0)
	for _, encounter := range apiResults.PokemonEncounters {
		cacheEntryVal = append(cacheEntryVal, encounter.Pokemon.Name)
	}
	pokecache.GetPokedexCacheInstance().Add(url, cacheEntryVal)
	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	printStringSlice(cacheEntryVal)
}

func PokedexCatchPokemon(name string) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	val, ok := checkAndPrintFromMem(url, false)
	var baseExp int
	var successChanceInStringSlice []string
	var successChanceInFloat float64
	if !ok {
		body, err := callAPI(url)
		if err != nil {
			fmt.Printf("Error when call API: %v\n", err)
			return
		}
		var apiResults types.PokedexPokemonAPI
		err = json.Unmarshal(body, &apiResults)
		if err != nil {
			fmt.Printf("Error when unmarshal: %v\n", err)
			return
		}
		baseExp = apiResults.BaseExperience //int
		successChanceInStringSlice, successChanceInFloat = calcCatchSuccessRate(baseExp)
		pokecache.GetPokedexCacheInstance().Add(url, successChanceInStringSlice)
	} else {
		successChanceInFloat, _ = strconv.ParseFloat(val[0], 64) // string to float
	}
	catchPokemonAndPrint(successChanceInFloat, name)
}

func calcCatchSuccessRate(sBase int) ([]string, float64) {
	maxBase := 608 * 1.1 //Blissey's base experience 608, hightest amoung all pokemons, ensure 10% for this as a baseline
	fRate := 1 - (float64(sBase) / maxBase)
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
	printStringSlice(printSlice)
}
