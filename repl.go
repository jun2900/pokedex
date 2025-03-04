package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jun2900/pokedexcli/internal"
)

func startRepl() {
	config := &config{
		NextURL: "https://pokeapi.co/api/v2/location-area/",
		Cache:   internal.NewCache(5 * time.Second),
		Pokedex: make(map[string]Pokemon),
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		extraCommand := ""
		if len(words) > 1 {
			extraCommand = words[1]
		}

		command, exists := config.getCommands(extraCommand)[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	NextURL string
	PrevURL *string
	Cache   *internal.Cache
	Pokedex map[string]Pokemon
}

func (c *config) getCommands(input string) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    c.commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays a map of the Pokedex",
			callback:    c.mapPokedexForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays a map of the Pokedex backward",
			callback:    c.mapPokedexBackward,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemon in a certain area",
			callback: func() error {
				return c.commandExplore(input)
			},
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback: func() error {
				return c.commandCatch(input)
			},
		},
		"inspect": {
			name:        "inspect",
			description: "Displays information about a pokemon",
			callback: func() error {
				return c.commandInspect(input)
			},
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays your pokedex",
			callback: func() error {
				return c.commandPokedex()
			},
		},
	}
}
