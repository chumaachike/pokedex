package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/chumaachike/pokedexcli/internal/pokecache"
)

func canCatch(baseExp int) bool {
	maxBaseExp := 600.0
	exp := float64(baseExp)

	// Higher baseExp → lower chance
	chance := 1.0 - (exp / maxBaseExp)

	// Clamp chance between 5% and 90%
	if chance < 0.05 {
		chance = 0.05
	}
	if chance > 0.90 {
		chance = 0.90
	}

	roll := rand.Float64() // 0.0–1.0
	return roll < chance
}

func fetchAndPrintLocations(cfg *config, startURL *url.URL) error {
	cache := pokecache.NewCache(10 * time.Second)
	var u *url.URL
	var locations Locations
	if startURL != nil {
		u = startURL
	} else {
		u, _ = url.Parse("https://pokeapi.co/api/v2/location-area/")
	}
	val, ok := cache.Get(u.String())

	if !ok {
		res, err := http.Get(u.String())
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return errors.New("bad status code")
		}
		if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
			return err
		}
		data, err := json.Marshal(locations)
		if err != nil {
			return err
		}
		cache.Add(u.String(), data)

	} else {
		if err := json.Unmarshal(val, &locations); err != nil {
			return err
		}
	}
	// update cfg
	if locations.Next != "" {
		cfg.next, _ = url.Parse(locations.Next)
	}
	if locations.Previous != "" {
		cfg.previous, _ = url.Parse(locations.Previous)
	}

	// print names
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
