package main

import (
	"fmt"
	"time"
)

/* Explication_Du_Programme :
Cer Programme contient les fonctions des différents menu et également la fonction qui permet de quitter le jeu et de revenir au menu in game
c'est ce programme qui fait appel au autre programme dans une suite logique contextuel
*/

func Menu_Start(Pointeur_joueurAsPersonnage *bool) { //affiche le menu de depart
	var input string //variable qui attend une action de l'utilisateur pour commencer une nouvelle partie

	fmt.Println("Bonjour jeune aventurier, vous venez de commencer une nouvelle avanture\n") //affiche le menu de depart et attend une action de l'utilisateur
	fmt.Println("Pour commencer une nouvelle partie vous devrer crée votre héro !\n")
	fmt.Println("Appuyer sur la touche n'importe qu'elle touche pour commencer\n")
	fmt.Scanln(&input)
	if input != " " {
		clear()
		fmt.Println("Création de votre personnage")
		clear()
		Creation_Personnage(Pointeur_joueurAsPersonnage)
		*Pointeur_joueurAsPersonnage = true
		fmt.Println("Votre personnage a été crée avec succès§§ Champion\n")
		Menu_In_Game(Pointeur_joueurAsPersonnage)
	}
}

func Menu_In_Game(Pointeur_joueurAsPersonnage *bool) { //affiche le menu in game
	var input_Menu rune // variable qui permet de naviger dans les sous menus
	if *Pointeur_joueurAsPersonnage == false {
		Menu_Start(Pointeur_joueurAsPersonnage)
	}
	clear()                                                     //clear le terminal
	fmt.Println("1. Afficher les informations du personnage\n") //affiche le menu
	fmt.Println("2. aller a l'invantaire\n")
	fmt.Println("3. aller au Shop\n")
	fmt.Println("4. Quitter le jeu\n")
	fmt.Scanf("%c", &input_Menu) //variable qui attend une action de l'utilisateur pour aller dans un sous menu
	var input_Menu2 string       //variable qui attend une action de l'utilisateur pour aller retournere au menu in game
	switch input_Menu {          //attend une action de l'utilisateur
	case '1': //si l'utilisateur appuie sur 1, affiche les informations du personnage
		clear()
		fmt.Println("fonction n'est pas encore disponible\n") // a faire
		Re_Menu_In_Game(Pointeur_joueurAsPersonnage)
	case '2': //si l'utilisateur appuie sur 2, affiche l'invantaire
		clear()
		//accessInventory()
		fmt.Println("Appuyer sur la touche n'importe qu'elle touche pour revenir au menu\n")
		fmt.Scanln(&input_Menu2)
		if input_Menu2 != " " {
			Menu_In_Game(Pointeur_joueurAsPersonnage)
		}

	case '3': // affiche le shop
		clear()
		fmt.Println("fonction n'est pas encore disponible\n") // a faire
		Re_Menu_In_Game(Pointeur_joueurAsPersonnage)
	case '4': // forgeron
		clear()
		fmt.Println("fonction n'est pas encore disponible\n") // a faire
		Re_Menu_In_Game(Pointeur_joueurAsPersonnage)

	case '5': // quitter le jeu
		clear()
		Quitter_Game() // fonction permet de mettre fin aux programme

	default: //si l'utilisateur appuie sur une autre touche, affiche une erreur
		Menu_In_Game(Pointeur_joueurAsPersonnage)
	}
}

func Quitter_Game() { //fonction qui quitte le jeu
	clear()
	fmt.Println("Vous avez quitter le jeu")
}

func Re_Menu_In_Game(Pointeur_joueurAsPersonnage *bool) { //fonction qui permet de revenir au menu in game
	fmt.Println("Retour dans le menu principal dans 2 secondes\n")
	time.Sleep(2 * time.Second)
	clear()
	Menu_In_Game(Pointeur_joueurAsPersonnage)
}
