package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jh123x/pokedex/internal/command"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommands map[string]cliCommand

func main() {
	cliCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    command.CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCmd,
		},
		"map": {
			name:        "map",
			description: "Locations",
			callback:    command.GetPokedexMapGen(true),
		},
		"mapb": {
			name:        "map",
			description: "Locations",
			callback:    command.GetPokedexMapGen(false),
		},
	}

	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		input, err := scanner.ReadString('\n')
		if err != nil {
			panic(err)
		}

		inputValues := cleanInput(input)
		if len(inputValues) == 0 {
			continue
		}

		command := strings.ToLower(inputValues[0])
		cmd, ok := cliCommands[command]
		if !ok {
			fmt.Printf("command not found: %s\n", command)
			continue
		}
		if err := cmd.callback(); err != nil {
			panic(err)
		}
	}
}

func helpCmd() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("")
	for _, v := range cliCommands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

func cleanInput(text string) []string {
	values := make([]string, 0)
	for _, word := range strings.Split(text, " ") {
		if word == "" {
			continue
		}

		values = append(values, strings.Trim(word, " \n"))
	}

	return values
}
