package main

import (
	"fmt"
	"time"
)

func (perso1 *Personnage) Fight_La_Chose(Droit_POtion_gratuit bool) {
	var Round_numéro int = 1
	var LaChose_training Monstre = Init_LaChose("Monstre d'entrainement", 40, 40, 5)
	fmt.Println("--------------------")
	fmt.Println("Vous êtes en combat contre La Chose :")
	fmt.Println("Tour numéro: ", Round_numéro)
	fmt.Println("Êtes-vous prêt ? y/n")
	fmt.Println("--------------------")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "y":
		clear()
		var Tour_suivant bool
		Nombre_de_tour := 0

		if Nombre_de_tour == 0 {
			Tour_suivant = randomBool()
			Nombre_de_tour++
		}

		for perso1.PV_Actuels > 0 && LaChose_training.PV_Actuels > 0 {
			time.Sleep(2 * time.Second)
			if Tour_suivant == true {
				clear()
				fmt.Println("A votre tour :")
				fmt.Println("Tour numéro : ", Nombre_de_tour)
				time.Sleep(2 * time.Second)
				Menu_tour_perso1(perso1, Droit_POtion_gratuit, &LaChose_training)
				Tour_suivant = false
			} else if Tour_suivant == false {
				clear()
				fmt.Println("Tour de l'ennemi :")
				fmt.Println("Tour numéro : ", Nombre_de_tour)
				fmt.Println("--------------------")
				time.Sleep(2 * time.Second)
				perso1.Attaque_Chose(&LaChose_training)

				time.Sleep(2 * time.Second)
				// Mettez ici la logique pour l'attaque de l'ennemi
				Tour_suivant = true
				Nombre_de_tour++
			}
		}
	case "n":
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	}
	if LaChose_training.PV_Actuels <= 0 {
		Dead(perso1)
	} else if perso1.PV_Actuels <= 0 {
		fmt.Println("Vous avez gagné le combat !")
		fmt.Println("comme c'est un entrainement vous ne gagnez pas de récompense, cheh!!")
		fmt.Println("--------------------")
		time.Sleep(2 * time.Second)
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	}
}
