package main

import "fmt"

func Forgeron(perso1 *Personnage, Droit_POtion_gratuit bool ) {
	var input_Menu_2 string
	fmt.Println("Bienvenue dans ma forge, veut-tu fabrique un Ityem ?")
	fmt.Println("-------------------------")
	fmt.Println("1. Fabriquer un chapeau d'aventurier ")
	fmt.Println("-------------------------")
	fmt.Println("2. Fabriquer Tunique de l'avanturier")
	fmt.Println("-------------------------")
	fmt.Println("3. Fabriquer Boittes de l'aventurier")
	fmt.Println("-------------------------")
	fmt.Println("4. Retour Menu Principal")
	fmt.Scanln(&input_Menu_2)
	switch input_Menu_2 {
	
	case "1":
		clear()
	case "2":
		clear()

	case "3":
		clear()
		
	case "4":
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	default:
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	}
}

