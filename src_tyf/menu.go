package main

import (
	"fmt"
	"time"
)

func Menu_Start(Pointeur_joueurAsPersonnage *bool) { //affiche le menu de depart
	var input string //variable qui attend une action de l'utilisateur pour commencer une nouvelle partie

	fmt.Println("appuyer sur une touche, pour commencer une nouvelle partie") //affiche le menu de depart et attend une action de l'utilisateur
	fmt.Scanln(&input)
	if input != " " {
		Menu_In_Game(Pointeur_joueurAsPersonnage)
	}
}

func Menu_In_Game(Pointeur_joueurAsPersonnage *bool) { //affiche le menu in game
	var input_Menu rune                                         // variable qui permet de naviger dans les sous menus                                            // variable qui permet de retourner aux menu in game
	clear()                                                     //clear le terminal
	fmt.Println("1. Afficher les informations du personnage\n") //affiche le menu
	fmt.Println("2. Accèder a votre Invantaire\n")
	fmt.Println("3. Crée votre personnage\n")
	fmt.Println("4. Quitter le jeu\n")
	fmt.Scanf("%c", &input_Menu)
	switch input_Menu { //attend une action de l'utilisateur
	case '1': //si l'utilisateur appuie sur 1, affiche les informations du personnage
		clear()
		if *Pointeur_joueurAsPersonnage == true {
			clear()
			fmt.Println("fonction n'est pas encore disponible\n") // a faire
			Re_Menu_In_Game(Pointeur_joueurAsPersonnage)
		} else if *Pointeur_joueurAsPersonnage == false {
			clear()
			fmt.Println("tu as pas de personnage!!!\n")
			fmt.Println("veuiller crée votre personnage pour commencer votre aventure !!\n")
			Re_Menu_In_Game(Pointeur_joueurAsPersonnage)

		}

	case '2': //si l'utilisateur appuie sur 2, affiche l'invantaire
		clear()
		//Invantaire()
		fmt.Println("vous avez pas d'invntaire\n")
		Re_Menu_In_Game(Pointeur_joueurAsPersonnage)

	case '3': //si l'utilisateur appuie sur 3, création du personnage
		clear()
		fmt.Println("Crée votre personnage")
		clear()
		Creation_Personnage()
		*Pointeur_joueurAsPersonnage = true
		fmt.Println("Votre personnage a été crée avec succès§§ Champion\n")
		Re_Menu_In_Game(Pointeur_joueurAsPersonnage)
	case '4': //si l'utilisateur appuie sur 4, quitte le jeu
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
