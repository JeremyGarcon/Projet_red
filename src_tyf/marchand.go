package main

import "fmt"

type Marchand struct {
	Inventaire  []Inventaire

}




for {
	fmt.Println("Menu:")
	fmt.Println("1. Utiliser une Potion de poison")
	fmt.Println("2. Acheter une Potion de poison auprès du marchand")
	fmt.Println("3. Quitter")

	var choix int
	fmt.Print("Choisissez une option: ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		// Utilisation de la Potion de poison
		poisonPot(personnage)
	case 2:
		// Achat d'une Potion de poison auprès du marchand
		acheterPotion(personnage, marchand)
	case 3:
		fmt.Println("Au revoir!")
		return // Quitter le programme
	default:
		fmt.Println("Option invalide. Veuillez choisir une option valide.")
	}
}