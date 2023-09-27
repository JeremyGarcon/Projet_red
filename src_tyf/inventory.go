package main

import (
	"fmt"
)

type Item struct { // Structure item
	Name string
	Price int
}

type Inventaire struct { // Structure inventaire
	Element Item
}

func (p *Personnage) accessInventory() {
	fmt.Println("Inventaire :")
	for _, item := range p.Inventaire {
		fmt.Println(" - ", item.Element.Name)
	}
}
