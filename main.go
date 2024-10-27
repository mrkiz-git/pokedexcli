package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mrkiz-git/pokedexcli/pokeapi"
)

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "Exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "Map",
			Description: "Displays the names of 20 location",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "Map Back",
			Description: "Displays the previous 20 locations",
			Callback:    CommandMapBack,
		},
	}
}

func CommandHelp(*CliConfig) error {
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for key, value := range commands {
		fmt.Printf("%s :%s\n", key, value.Description)
	}

	return nil
}

func CommandMap(config *CliConfig) error {
	log.Println("Fetching locations...")

	resp, err := config.pokeapiClient.ListLocationAreas(config.NextLocationUrl)
	if err != nil {
		log.Printf("Error fetching locations: %v", err)
		return err
	}

	config.NextLocationUrl = resp.Next
	config.PrevLocationUrl = resp.Previous

	log.Printf("Found %d locations", len(resp.Results))

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandMapBack(config *CliConfig) error {
	if config.PrevLocationUrl == nil {
		log.Println("No previous page available")
		return fmt.Errorf("you are on the first page")
	}

	log.Println("Fetching previous locations...")

	resp, err := config.pokeapiClient.ListLocationAreas(config.PrevLocationUrl)
	if err != nil {
		log.Printf("Error fetching previous locations: %v", err)
		return err
	}

	config.NextLocationUrl = resp.Next
	config.PrevLocationUrl = resp.Previous

	log.Printf("Found %d locations", len(resp.Results))

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandExit(config *CliConfig) error {
	os.Exit(0)
	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := GetCommands()

	cliConfig := &CliConfig{
		NextLocationUrl: nil,
		PrevLocationUrl: nil,
		pokeapiClient:   pokeapi.New(),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		if command, exists := commands[input]; exists {
			err := command.Callback(cliConfig)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}

func main() {
	log.SetPrefix("Pokedex: ")
	startRepl()
}
