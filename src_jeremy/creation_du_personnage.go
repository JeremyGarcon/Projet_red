package main

import (
	"fmt"
	"time"
)

type Personnage struct {
	Nom        string
	Classe     string
	Niveau     int
	PV_Max     int
	PV_Actuels int
	Sort       string
}

func charCreation() Personnage {
	var Nom string
	var Classe string
	clear()
	fmt.Println("Character creation in progress...")
	time.Sleep(1 * time.Second)
	clear()
	fmt.Print("Please enter a name: ")
	fmt.Scanln(&Nom)
	clear()
	fmt.Print("Please choose a category between Human, Elf, and Dwarf: ")
	fmt.Scanln(&Classe)
	fmt.Print("\n")

	var PV_Max int
	var Sort string

	switch Classe {
	case "Human":
		PV_Max = 100
		Sort = "Punch"
	case "Elf":
		PV_Max = 80
		Sort = "Swift Strike"
	case "Dwarf":
		PV_Max = 120
		Sort = "Hard Punch"
	default:
		clear()
		return charCreation() // Rappel récursif si la classe est invalide
	}

	joueur := Personnage{
		Nom:        Nom,
		Classe:     Classe,
		Niveau:     1,
		PV_Max:     PV_Max,
		PV_Actuels: PV_Max / 2,
		Sort:       Sort,
	}

	return joueur
}

func Creation_Personnage() {
	personnage := charCreation()

	fmt.Println("Your character has been successfully created!")
	fmt.Printf("\n")
	fmt.Printf("Name: %s\n", personnage.Nom)
	fmt.Printf("Category: %s\n", personnage.Classe)
	fmt.Printf("Level: %d\n", personnage.Niveau)
	fmt.Printf("Pv Max: %d\n", personnage.PV_Max)
	fmt.Printf("Pv : %d\n", personnage.PV_Actuels)
	fmt.Printf("Spell: %s\n", personnage.Sort)
}
