package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CliCommand represents a command in the CLI application
type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

// GetCommands returns a map of all available CLI commands
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
			Description: "Exit the Pokedex",
			Callback:    nil,
		},
	}
}

// CommandHelp displays help information
func CommandHelp() error {
	// Implementation of help command
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for key, value := range commands {
		fmt.Printf("%s :%s\n", key, value.Description)
	}

	return nil
}

// CommandExit handles the exit command
func CommandExit() error {
	// Implementation of exit command
	os.Exit(0)
	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := GetCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		if command, exists := commands[input]; exists {
			err := command.Callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}

func main() {
	//create cli scanner
	startRepl()
}
