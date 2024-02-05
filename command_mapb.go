package main

import (
	"fmt"

	"github.com/carlcrede/pokedex-cli/internal"
)

func commandMapb(cfg *config) error {
	locationAreas, err := internal.GetPreviousLocationAreas(cfg.Previous, cfg.cache)
	if err != nil {
		return err
	}
	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous
	return nil
}
