package main

import "fmt"



/* 
Explication du programme :
Ce programme définis les fonctions qui affichent les informations du personnage dans le menu du jeu
*/ 

func Display_Info_Perso1(perso1 *Personnage, Droit_POtion_gratuit bool ) { // fonction qui affiche les informations du personnage dans le case 2 de la foncti_on Menu_In_Game
	clear()
	fmt.Println("Information du joueur: ")
	fmt.Println("-------------------------")
	fmt.Printf("Name: %s \n", perso1.Nom)
	fmt.Printf("Classe: %s \n", perso1.Classe)
	fmt.Printf("Niveau: %d \n", perso1.Niveau)         // Utilisez %d pour les entiers
	fmt.Printf("PV Max: %d \n", perso1.PV_Max)         // Utilisez %d pour les entiers
	fmt.Printf("PV Actuels: %d \n", perso1.PV_Actuels) // Utilisez %d pour les entiers
	fmt.Printf("Argent: %d \n", perso1.Argent)         // Utilisez %d pour les entiers
	fmt.Println("-------------------------")
	fmt.Println("Appuyer sur une touche pour revenir au menu")
	var input_Menu_1 string
	fmt.Scanln(&input_Menu_1)
	if input_Menu_1 != " " {
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	}

}
func Display_Info_Sorts(perso1 *Personnage, Droit_POtion_gratuit bool) {
	clear()
	fmt.Println("Information des sort du joueur ")
	fmt.Println("-------------------------")
	for i := 0; i < len(perso1.Skill); i++ {
		fmt.Printf("%d. %s \n", i, perso1.Skill[i].Name)
	}

	fmt.Println("--------------------")
	fmt.Println("Appuyer sur une touche pour revenir au menu")
	var input_Menu_1 string
	fmt.Scanln(&input_Menu_1)
	if input_Menu_1 != " " {
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	}
}
