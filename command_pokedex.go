package main

import "fmt"

func commandPokedex(cfg *config) error {
	if len(cfg.caugthPokemon) < 1 {
		fmt.Println("You haven't caught any pokemon yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caugthPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
