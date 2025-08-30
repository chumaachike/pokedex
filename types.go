package main

import (
	"net/url"
)

type config struct {
	next     *url.URL
	previous *url.URL
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type Locations struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Resource `json:"results"`
}

type Resource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             Resource              `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod EncounterMethod `json:"encounter_method"`
	VersionDetails  []VersionDetail `json:"version_details"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionDetail struct {
	Rate    int      `json:"rate"`
	Version Resource `json:"version"`
}

type Name struct {
	Language Resource `json:"language"`
	Name     string   `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        Resource               `json:"pokemon"`
	VersionDetails []PokemonVersionDetail `json:"version_details"`
}

type PokemonVersionDetail struct {
	EncounterDetails []Encounter `json:"encounter_details"`
	MaxChance        int         `json:"max_chance"`
	Version          Resource    `json:"version"`
}

type Encounter struct {
	Chance          int        `json:"chance"`
	MinLevel        int        `json:"min_level"`
	MaxLevel        int        `json:"max_level"`
	ConditionValues []Resource `json:"condition_values"`
	Method          Resource   `json:"method"`
}

type Pokemon struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	BaseExperience int              `json:"base_experience"`
	Height         int              `json:"height"`
	IsDefault      bool             `json:"is_default"`
	Order          int              `json:"order"`
	Weight         int              `json:"weight"`
	Abilities      []PokemonAbility `json:"abilities"`
	Forms          []Resource       `json:"forms"`
	//GameIndices          []VersionGameIndex `json:"game_indices"`
	// HeldItems            []PokemonHeldItem  `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	// Moves                []PokemonMove      `json:"moves"`
	//PastTypes            []PokemonTypePast  `json:"past_types"`
	Sprites PokemonSprites `json:"sprites"`
	Cries   PokemonCries   `json:"cries"`
	Species Resource       `json:"species"`
	Stats   []PokemonStat  `json:"stats"`
	Types   []PokemonType  `json:"types"`
}

type PokemonAbility struct {
	IsHidden bool     `json:"is_hidden"`
	Slot     int      `json:"slot"`
	Ability  Resource `json:"ability"`
}

type PokemonType struct {
	Slot int      `json:"slot"`
	Type Resource `json:"type"`
}

type PokemonStat struct {
	Stat     Resource `json:"stat"`
	Effort   int      `json:"effort"`
	BaseStat int      `json:"base_stat"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
