package main

import (
	"fmt"
	"time" )

type Personnage struct { // Structure pour le personnage
	Nom        string
	Classe     string
	Niveau     int8
	PV_Max     int16
	PV_Actuels int16
	Sort       string
	Inventaire []Inventaire
}

func charCreation() Personnage {  // Création du personnage
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

	var PV_Max int16
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
		fmt.Print("Invalid Category. Please choose again.")
		return charCreation() // Rappel récursif si la classe est invalide
	}

	inventoryInitiale := []Inventaire{ // Inventaire du joueur
		{Element: Item{Name: "Epée"}},
	}

	joueur := Personnage{
		Nom:        Nom,
		Classe:     Classe,
		Niveau:     1,
		PV_Max:     PV_Max,
		PV_Actuels: PV_Max / 2,
		Sort:       Sort,
		Inventaire: inventoryInitiale,
	}
	return joueur
}

func Creation_Personnage() {

	personnage := charCreation()

	fmt.Println("Your character has been successfully created!") // Affichage des informations dans le terminal
	fmt.Printf("\n")
	fmt.Printf("Name: %s\n", personnage.Nom)
	fmt.Printf("Category: %s\n", personnage.Classe)
	fmt.Printf("Level: %d\n", personnage.Niveau)
	fmt.Printf("Pv Max: %d\n", personnage.PV_Max)
	fmt.Printf("Pv : %d\n", personnage.PV_Actuels)
	fmt.Printf("Spell: %s\n", personnage.Sort)
	personnage.accessInventory()
}
