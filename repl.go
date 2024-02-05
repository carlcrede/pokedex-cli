package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/carlcrede/pokedex-cli/internal"
	"github.com/carlcrede/pokedex-cli/internal/pokecache"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(10 * time.Second)
	cfg := &config{
		cache:         cache,
		caugthPokemon: make(map[string]internal.PokemonResponse),
	}
	for {
		fmt.Printf("pokedex 💻 > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		// add params to cfg if exists
		if len(words) > 1 {
			cfg.Params = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println("Error executing command: ", err)
			}
			continue
		} else {
			fmt.Println("Invalid command: ", commandName)
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type config struct {
	Next          string
	Previous      string
	Params        []string
	cache         *pokecache.Cache
	caugthPokemon map[string]internal.PokemonResponse
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "(map back) Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "See a list of all the Pokémon in the a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try to catch a pokemon! If successfull, its added to your Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "See details about a pokemon you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists names of all the pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}
