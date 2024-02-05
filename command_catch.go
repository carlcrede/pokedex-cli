package main

import (
	"fmt"
	"math/rand"

	"github.com/carlcrede/pokedex-cli/internal"
)

func commandCatch(cfg *config) error {
	pokemon_name := cfg.Params[0]
	pokemon, err := internal.GetPokemon(pokemon_name, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)

	randomFloat := rand.Float32()

	user_catch_chance := calculateUserCatchChance(pokemon.BaseExperience)

	if randomFloat <= user_catch_chance {
		fmt.Printf("%s was caught!\n", pokemon_name)
		cfg.caugthPokemon[pokemon_name] = *pokemon
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%s escaped!\n", pokemon_name)
	}

	return nil
}

func calculateUserCatchChance(pokemon_base_exp int) float32 {
	if pokemon_base_exp >= 250 {
		return .1
	}
	if pokemon_base_exp >= 200 {
		return .15
	}
	if pokemon_base_exp >= 150 {
		return .2
	}
	if pokemon_base_exp >= 100 {
		return .25
	}
	if pokemon_base_exp >= 50 {
		return .3
	}
	return .5
}
