package main

import (
	"fmt"
	"time"
)

/*
	Explication_Du_Programme :

Ce Programme définis les structures des items et de l'inventaire du joueur, puis crée les instance de chaque structure
*/

// structure d'item :
type Item struct { //structure des items
	Name        string //SSID de l'item
	id          int    // numéro de l'item
	Value       int    // Valeur de soin ou d'attaque ex: potion = 100 donne 100 pv et a contrario épée = -50 enlève 50 pv a l'adversarie
	description string // description de l'item
}

// structure des différents Items :
var Epee_En_Bois = Item{ // item de base

	Name:        "Eppée en bois",
	id:          1,
	Value:       5,
	description: "Une eppée en bois, pas très résistant mais aide a se défendre",
}
var Potion_De_Soin_Nv0 = Item{ // item de base

	Name:        "Potion de soin Nv0",
	id:          2,
	Value:       7,
	description: "Une potion de soin Nv0, ne regénère que très peu de pv",
}
var Potion_De_Poison_Nv0 = Item{ // item de base

	Name:        "Potion de Poison Nv0",
	id:          3,
	Value:       -7,
	description: "Une potion de Poison Nv0, ne fait que très peu de dégat aux monstres",
}

// Structure de l'inventaire :
type Personnage_intventory struct { // Structure de l'inventaire
	Inventaire []Item
}

// Structure pour le personnage :
type Personnage struct { // Structure pour le personnage
	Nom        string
	Classe     string
	Niveau     int8
	PV_Max     int16
	PV_Actuels int16
	Sort       string
	Inventaire []Item
}

func charCreation(Pointeur_joueurAsPersonnage *bool) Personnage { // Création du personnage
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

	switch Classe { // Choix de la classe ce qui vas influencer le contenue des variables PV_Max et Sort du joueur
	case "Human":
		PV_Max = 100 // Initialisation des variables en fonction de la classe
		Sort = "Punch"
	case "Elf":
		PV_Max = 80 // Initialisation des variables en fonction de la classe
		Sort = "Swift Strike"
	case "Dwarf":
		PV_Max = 120 // Initialisation des variables en fonction de la classe
		Sort = "Hard Punch"
	default:
		clear()
		fmt.Print("Invalid Category. Please choose again.")
		return charCreation(Pointeur_joueurAsPersonnage) // Rappel récursif si la classe est invalide
	}

	joueur := Personnage{
		Nom:        Nom, // Initialisation des variables en fonction de la classe
		Classe:     Classe,
		Niveau:     1,
		PV_Max:     PV_Max,
		PV_Actuels: PV_Max / 2,
		Sort:       Sort,
		Inventaire: []Item{},
	}
	return joueur
}

func Creation_Personnage(Pointeur_joueurAsPersonnage *bool) {

	personnage := charCreation(Pointeur_joueurAsPersonnage)

	fmt.Println("Your character has been successfully created!") // Affichage des informations dans le terminal
	fmt.Printf("\n")
	fmt.Printf("Name: %s\n", personnage.Nom)
	fmt.Printf("Category: %s\n", personnage.Classe)
	fmt.Printf("Level: %d\n", personnage.Niveau)
	fmt.Printf("Pv Max: %d\n", personnage.PV_Max)
	fmt.Printf("Pv : %d\n", personnage.PV_Actuels)
	fmt.Printf("Spell: %s\n", personnage.Sort)
}

func Create_Inventory() Personnage_intventory { //création de l'inventaire du joueur

	joueur := Personnage_intventory{

		Inventaire: []Item{Epee_En_Bois, Potion_De_Soin_Nv0, Potion_De_Poison_Nv0},
	}
	return joueur
}

/*
func display_Info() { // affiche les informations du personnage
	fmt.Printf("Name: %s\n", personnage.Nom)
	fmt.Printf("Category: %s\n", personnage.Classe)
	fmt.Printf("Level: %d\n", personnage.Niveau)
	fmt.Printf("Pv Max: %d\n", personnage.PV_Max)
	fmt.Printf("Pv : %d\n", personnage.PV_Actuels)
	fmt.Printf("Spell: %s\n", personnage.Sort)
}
*/
