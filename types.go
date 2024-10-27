package main

import (
	"github.com/mrkiz-git/pokedexcli/pokeapi" // Import your package
)

// CliConfig represents the needed data for this session
type CliConfig struct {
	NextLocationUrl *string
	PrevLocationUrl *string
	pokeapiClient   *pokeapi.Client
}

// CliCommand represents a command in the CLI application
type CliCommand struct {
	Name        string
	Description string
	Callback    func(*CliConfig) error
}
