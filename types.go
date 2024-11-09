package main

import (
	"github.com/mrkiz-git/pokedexcli/pokeapi"
)

// CliConfig represents the needed data for this session
type CliConfig struct {
	NextLocationUrl *string
	PrevLocationUrl *string
	pokeapiClient   *pokeapi.Client
	pokedex         map[string]pokeapi.Pokemon
}

// CliCommand represents a command in the CLI application
type CliCommand struct {
	Name        string
	Description string
	Callback    func(*CliConfig, []string) error
}
