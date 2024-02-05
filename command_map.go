package main

import (
	"fmt"

	"github.com/carlcrede/pokedex-cli/internal"
)

func commandMap(cfg *config) error {
	locationAreas, err := internal.GetNextLocationAreas(cfg.Next, cfg.cache)
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
