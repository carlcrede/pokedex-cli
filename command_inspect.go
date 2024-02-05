package main

import (
	"fmt"
)

func commandInspect(cfg *config) error {
	name := cfg.Params[0]
	value, ok := cfg.caugthPokemon[name]
	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", value.Name)
	fmt.Printf("Height: %d\n", value.Height)
	fmt.Printf("Weight: %d\n", value.Weight)

	hp := value.Stats[0].BaseStat
	attack := value.Stats[1].BaseStat
	defence := value.Stats[2].BaseStat
	specialAttack := value.Stats[3].BaseStat
	specialDefence := value.Stats[4].BaseStat
	speed := value.Stats[5].BaseStat

	fmt.Println("Stats:")
	fmt.Printf("  -hp: %d\n  -attack: %d\n  -defence: %d\n  -special-attack: %d\n  -special-defence: %d\n  -speed: %d\n", hp, attack, defence, specialAttack, specialDefence, speed)
	fmt.Println("Types:")

	for _, t := range value.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
