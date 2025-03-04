package main

import "fmt"

func (c *config) commandPokedex() error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.Pokedex {
		fmt.Println(pokemon.Name)
	}
	return nil
}
