package main

import (
	"fmt"

	"github.com/carlcrede/pokedex-cli/internal"
)

func commandExplore(cfg *config) error {
	area_name := cfg.Params[0]
	locationArea, err := internal.GetLocationArea(area_name, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area_name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
