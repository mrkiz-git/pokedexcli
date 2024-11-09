package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/mrkiz-git/pokedexcli/pokeapi"
)

func GetCommands() map[string]CliCommand {
	log.Println("Initializing commands")
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
		"explore": {
			Name:        "Explore location",
			Description: "Displays all the Pokémon in a given areas",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "Catch Pockemon",
			Description: "Catch Pockemon",
			Callback:    CommandCatch,
		},
	}
}

func CommandHelp(config *CliConfig, args []string) error {

	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for key, value := range commands {
		fmt.Printf("%s :%s\n", key, value.Description)
	}

	return nil
}

func CommandMap(config *CliConfig, args []string) error {

	log.Println("Fetching locations...")

	resp, err := config.pokeapiClient.GetLocationAreas(config.NextLocationUrl)
	if err != nil {
		log.Printf("Error fetching locations: %v", err)
		return err
	}

	log.Printf("Next URL was: %v", config.NextLocationUrl)
	log.Printf("Previous URL was: %v", config.PrevLocationUrl)
	log.Printf("Setting Next URL to: %v", resp.Next)
	log.Printf("Setting Previous URL to: %v", resp.Previous)

	config.NextLocationUrl = resp.Next
	config.PrevLocationUrl = resp.Previous

	log.Printf("Found %d locations", len(resp.Results))

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandMapBack(config *CliConfig, args []string) error {
	log.Println("Executing mapb command")

	if config.PrevLocationUrl == nil {

		return fmt.Errorf("you are on the first page")
	}

	log.Printf("Using previous URL: %v", *config.PrevLocationUrl)

	resp, err := config.pokeapiClient.GetLocationAreas(config.PrevLocationUrl)
	if err != nil {
		log.Printf("Error fetching previous locations: %v", err)
		return err
	}

	log.Printf("Setting Next URL to: %v", resp.Next)
	log.Printf("Setting Previous URL to: %v", resp.Previous)

	config.NextLocationUrl = resp.Next
	config.PrevLocationUrl = resp.Previous

	log.Printf("Found %d locations", len(resp.Results))

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandExplore(config *CliConfig, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("wrong number of arguments provided")
	}
	log.Printf("Executing Explore Command for %s", args[0])

	resp, err := config.pokeapiClient.GetLocation(&args[0])
	if err != nil {
		log.Printf("Error geting location %v", err)
		return err
	}

	if len(resp.PokemonEncounters) < 1 {
		fmt.Println("No Pokemon found at this location")
		return nil
	}

	fmt.Printf("Exploring %s...", args[0])
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println("- ", encounter.Pokemon.Name)
	}
	return nil

}

func CommandCatch(config *CliConfig, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("wrong number of arguments provided\n")
	}

	resp, err := config.pokeapiClient.GetPockemon(&args[0])
	if err != nil {
		log.Printf("Error geting pockemon %v", err)
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s\n", args[0])
	randomNumber := rand.Intn(501)
	fmt.Printf("Base Experience: %d\n", resp.BaseExperience)
	fmt.Printf("Random Number: %d\n", randomNumber)

	if randomNumber >= resp.BaseExperience {
		config.pokedex[args[0]] = *resp
		fmt.Printf("%s was caught!\n", resp.Name)
		return nil
	} else {
		fmt.Printf("%s was escaped!\n", resp.Name)
		return nil
	}

}

func CommandExit(config *CliConfig, args []string) error {

	os.Exit(0)
	return nil
}

func startRepl() {
	log.Println("Starting REPL")
	scanner := bufio.NewScanner(os.Stdin)
	commands := GetCommands()

	log.Println("Initializing CLI config")
	cliConfig := &CliConfig{
		NextLocationUrl: nil,
		PrevLocationUrl: nil,
		pokeapiClient:   pokeapi.New(),
		pokedex:         make(map[string]pokeapi.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		words := strings.Fields(input)
		if len(words) == 0 {
			log.Println("Empty input received")
			continue
		}

		commandName := words[0]
		args := words[1:]

		if command, exists := commands[commandName]; exists {

			err := command.Callback(cliConfig, args)
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
	log.SetFlags(log.Ltime | log.Ldate) // Add timestamps to logs
	log.Println("Starting Pokedex CLI")
	startRepl()
}
