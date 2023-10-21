package main

import "fmt"

func Menu_Choise_fight_training(perso1 *Personnage,Droit_POtion_gratuit bool) {
	clear()
	fmt.Println("--------------------")
	fmt.Println("Choisissez votre adversaire: ")
	fmt.Println("1. Fight: La Chose")
	fmt.Println("2. Fight: Le Gobelin")
	fmt.Println("3. retour Menu")
	fmt.Println("--------------------")
	fmt.Println("Appuyer sur une touche compris entre 1 et 3 pour naviguer dans le menu")
	fmt.Println("--------------------")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 1:
		clear()
		perso1.Fight_La_Chose(Droit_POtion_gratuit)
	case 2:
		clear() // a faire 
	case 3:
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
}
}
