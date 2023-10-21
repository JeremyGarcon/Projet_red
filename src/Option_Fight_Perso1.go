package main

import (
	"fmt"
	"time"
)

func Menu_tour_perso1(perso1 *Personnage, Droit_POtion_gratuit bool, LaChose_training *Monstre) {
	fmt.Println("--------------------")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser un item")
	fmt.Println("3. Utiliser un sort")
	fmt.Println("4. Fuite")
	fmt.Println("--------------------")

	var input_Menu string
	fmt.Print("Sélectionnez une option : ")
	fmt.Scanln(&input_Menu)

	switch input_Menu {
	case "1":
		Attaque_perso1(perso1, Droit_POtion_gratuit, LaChose_training)
	case "2":
		clear()
		perso1.use_item(LaChose_training)
	case "3":
		fmt.Println("fonction pas dispo pour le moment")
		time.Sleep(2 * time.Second)
		clear()
		Menu_tour_perso1(perso1, Droit_POtion_gratuit, LaChose_training)
	case "4":
		fuite(perso1, Droit_POtion_gratuit)
	}
}

func fuite(perso1 *Personnage, Droit_POtion_gratuit bool) {
	clear()
	fmt.Println("Vous avez fui le combat")
	time.Sleep(2 * time.Second)
	Menu_Choise_fight_training(perso1, Droit_POtion_gratuit)
}

func (perso1 *Personnage) use_item(LaChose_training *Monstre) { //fonction pour utiliser un item
	fmt.Println("Vous avez utilisé un item : ")
	fmt.Println("-------------------------")
	for i := 0; i < len(perso1.Inventaire); i++ { // Afficher les items de l'inventaire
		fmt.Printf("%d. %s \n", i, perso1.Inventaire[i].Name)
	}
	fmt.Println("-------------------------")
	var input_Item int
	fmt.Print("Sélectionnez un item (seules les Potions sont utilisables de manière active et à usage unique) : ")
	fmt.Scanln(&input_Item)

	// Vérifier si l'indice de l'item est valide
	if input_Item >= 0 && input_Item < len(perso1.Inventaire) { // Si l'indice est valide
		selectedItem := perso1.Inventaire[input_Item].Name

		// Vérifier le nom de l'item sélectionné
		switch selectedItem { // en fonction du nom de l'item
		case "Potion de soin Nv0":
			// Code pour utiliser une potion de soin
			fmt.Println("Vous avez utilisé une Potion de soin Nv0.")
			perso1.PV_Actuels += 10
			// Supprimer l'item de l'inventaire après utilisation (si nécessaire)
			perso1.Inventaire = append(perso1.Inventaire[:input_Item], perso1.Inventaire[input_Item+1:]...)
		case "Potion de Poison Nv0":
			LaChose_training.PV_Actuels -= 10
			fmt.Println("Vous avez utilisé une Potion de Poison Nv0.")
			// Supprimer l'item de l'inventaire après utilisation (si nécessaire)
			perso1.Inventaire = append(perso1.Inventaire[:input_Item], perso1.Inventaire[input_Item+1:]...)
		default:
			clear()
			fmt.Println("Item non reconnu.")
			time.Sleep(2 * time.Second)
			perso1.use_item(LaChose_training) // Appeler la fonction à nouveau en cas d'erreur
		}
	} else {
		clear()
		fmt.Println("Indice d'item invalide.")
		time.Sleep(2 * time.Second)
		perso1.use_item(LaChose_training) // Appeler la fonction à nouveau en cas d'erreur
	}
}

func Attaque_perso1(perso1 *Personnage, Droit_POtion_gratuit bool, LaChose_training *Monstre) {
	fmt.Println("Vous avez attaqué l'ennemi")
	
}
