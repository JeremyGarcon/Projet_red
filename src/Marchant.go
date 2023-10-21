package main

import (
	"fmt"
	"time"
)

func Marchant(perso1 *Personnage, Droit_POtion_gratuit bool) { // fonction qui affiche le shop dans le case 4 de la fonction Menu_In_Game
	clear()
	var input_Menu_2 string
	fmt.Println("Bienvenue dans mon magasin")
	fmt.Println("-------------------------")
	fmt.Println("Comme tu me s'emble nouveaux je t'offre une potion de vie (appuyer sur 1 pour l'accepter)")
	fmt.Println("-------------------------")
	fmt.Println("Acheter une potion de vie (appuyer sur 2 pour l'acheter)(coûte 3 pièces d'or)")
	fmt.Println("-------------------------")
	fmt.Println("Acheter une potion de poison (appuyer sur 3 pour l'acheter)(coute 6 pièces d'or)")
	fmt.Println("-------------------------")
	fmt.Println("Acheter un livre de sort (appuyer sur 4 pour l'acheter)(cout 25 pièces d'or)")
	fmt.Println("-------------------------")
	fmt.Println("Augmenter la taille de votre invantaire (appuyer sur 5 pour acheter) coute 50 pièces d'or")
	fmt.Println("-------------------------")
	fmt.Println("appuyer sur 6 pour quitter le magasin")
	fmt.Println("-------------------------")
	fmt.Scanln(&input_Menu_2)
	switch input_Menu_2 {
	case "1": // ajoute une potion de vie dans l'inventaire du joueur pour les nouveaux joueur
		clear()
		if Droit_POtion_gratuit == true {
		AddItemToInventory(perso1, Potion_De_Soin_Nv0)
		Droit_POtion_gratuit = false
		} else if Droit_POtion_gratuit == false{
			fmt.Println("Vous n'avez pas le droit de prendre cette potion")
		}
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	case "2": // acheter une potion de vie
		clear()
		GoldCheck(perso1, 3)
		if GoldCheck(perso1, 3) == true {
			AddItemToInventory(perso1, Potion_De_Soin_Nv0)
			perso1.Argent = perso1.Argent - 3
		} else if GoldCheck(perso1, 3) == false {
			fmt.Println("Vous n'avez pas assez d'argent")
			time.Sleep(2 * time.Second)
			clear()
			Menu_In_Game(perso1, Droit_POtion_gratuit)
		}
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	case "3": //acheter une potion de poison
			clear()
			GoldCheck(perso1, 3)
			if GoldCheck(perso1, 3) == true {
				AddItemToInventory(perso1, Potion_De_Poison_Nv0)
				perso1.Argent = perso1.Argent - 3
			} else if GoldCheck(perso1, 3) == false {
				fmt.Println("Vous n'avez pas assez d'argent")
				time.Sleep(2 * time.Second)
				clear()
				Menu_In_Game(perso1, Droit_POtion_gratuit)
	}
	Menu_In_Game(perso1, Droit_POtion_gratuit)
	case "4": //acheter un livre de sort

	case "5": //augmenter la taille de votre invantaire
		clear()
		if GoldCheck(perso1, 30) == true {
		perso1.UpgradeInventory()
		time.Sleep(2 * time.Second)
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
		} else if GoldCheck(perso1, 30) == false {
			fmt.Println("Vous n'avez pas assez d'argent")
			time.Sleep(2 * time.Second)
			clear()
			Menu_In_Game(perso1, Droit_POtion_gratuit)
		}

	case "6": //retour au menu principale
		clear()
		Menu_In_Game(&Personnage{}, Droit_POtion_gratuit)
	}
}
