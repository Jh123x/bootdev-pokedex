package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jh123x/pokedex/internal/command"
	"github.com/Jh123x/pokedex/internal/consts"
)

type cliCommand struct {
	name        string
	description string
	callback    consts.Command
}

var (
	cliCommands map[string]cliCommand
)

func main() {
	playerInfo := &consts.PlayerInfo{
		CaughtPokemons: make(map[string]*consts.PokemonInspectInfo, 100),
	}
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
			description: "Next locations",
			callback:    command.GetPokedexMapGen(true),
		},
		"mapb": {
			name:        "map",
			description: "Previous location page",
			callback:    command.GetPokedexMapGen(false),
		},
		"explore": {
			name:        "explore",
			description: "Explore location to get pokemons",
			callback:    command.CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    command.CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Look at your caught pokemon",
			callback:    command.CommandInspect,
		},
	}

	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		input, err := scanner.ReadString('\n')
		if err != nil {
			panic(err)
		}

		command, args := cleanInput(input)
		if command == "" {
			continue
		}

		command = strings.ToLower(command)
		cmd, ok := cliCommands[command]
		if !ok {
			fmt.Printf("command not found: %s\n", command)
			continue
		}
		if err := cmd.callback(args, playerInfo); err != nil {
			panic(err)
		}
	}
}

func helpCmd(_ []string, _ *consts.PlayerInfo) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("")
	for _, v := range cliCommands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

func cleanInput(text string) (string, []string) {
	values := make([]string, 0)
	for _, word := range strings.Split(text, " ") {
		if word == "" {
			continue
		}

		values = append(values, strings.Trim(word, " \n"))
	}

	if len(values) == 0 {
		return "", nil
	}

	return values[0], values[1:]
}
