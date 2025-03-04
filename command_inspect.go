package main

import "fmt"

func (c *config) commandInspect(pokemonName string) error {
	if data, found := c.Pokedex[pokemonName]; found {
		fmt.Printf("Name: %s\n", data.Name)
		fmt.Printf("Height: %d\n", data.Height)
		fmt.Printf("Weight: %d\n", data.Weight)
		for _, stat := range data.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range data.Types {
			fmt.Println("  -", typeInfo.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
