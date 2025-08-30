package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/chumaachike/pokedexcli/internal/pokecache"
)

var userPokedex = map[string]Pokemon{}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no pokemon provided")
	}
	pokemon, ok := userPokedex[args[0]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}
	return nil
}
func commandPokedex(cfg *config, args []string) error {
	if len(userPokedex) == 0 {
		fmt.Println("No pokemon has been caught")
	} else {
		fmt.Println("Your Pokedex")
		for k := range userPokedex {
			fmt.Printf("  - %s\n", k)
		}
	}
	return nil
}

func commandCatch(cfg *config, args []string) error {
	cache := pokecache.NewCache(10 * time.Second)
	var pokemon Pokemon

	if len(args) == 0 {
		return errors.New("np pokemon provided")
	}
	pokemonName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	val, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return errors.New("bad status code")
		}
		if err := json.NewDecoder(res.Body).Decode(&pokemon); err != nil {
			return nil
		}
		data, err := json.Marshal(pokemon)
		if err != nil {
			return err
		}
		cache.Add(url, data)

	} else {
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return err
		}
	}
	fmt.Printf("Throwing a Pokeball at %s... \n", pokemon.Name)
	if canCatch(pokemon.BaseExperience) {
		fmt.Printf("%s was caught! \n", pokemon.Name)
		userPokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped! \n", pokemon.Name)
	}
	return nil
}
func commandExplore(cfg *config, args []string) error {
	cache := pokecache.NewCache((10 * time.Second))
	var locationArea LocationArea
	if len(args) == 0 {
		return errors.New("no location area provided")
	}
	locArea := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locArea)

	val, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return errors.New("bad status code")
		}

		if err := json.NewDecoder(res.Body).Decode(&locationArea); err != nil {
			return err
		}
		data, err := json.Marshal(locationArea)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	} else {
		if err := json.Unmarshal(val, &locationArea); err != nil {
			return err
		}
	}

	fmt.Println("Exploring pastoria-city-area...")
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounters := range locationArea.PokemonEncounters {
		fmt.Println(" - ", pokemonEncounters.Pokemon.Name)

	}

	return nil

}

func commandMap(cfg *config, args []string) error {
	return fetchAndPrintLocations(cfg, cfg.next)
}

func commandMapb(cfg *config, args []string) error {
	return fetchAndPrintLocations(cfg, cfg.previous)
}

var registerMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Displays the next 20 location areas in the Pokemon world",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the previous 20 location areas in the Pokemon world",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Lists all the Pokemon located in an area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Takes the name of a Pokemon and catches it with some probability",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspects the characteristics of a Pokemon if it has been caught",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Prints the names of all the Pokemon in user's pokedex",
		callback:    commandPokedex,
	},
}

func init() {
	registerMap["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func(cfg *config, args []string) error {
			fmt.Println("Welcome to the Pokedex!")
			fmt.Println("Usage:")
			for _, v := range registerMap {
				fmt.Printf("  %s: %s\n", v.name, v.description)
			}
			return nil
		},
	}
}
