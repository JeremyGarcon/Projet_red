package main

// Structure du personnage

import (
	"fmt"
	"time"
)

func main() { // fonction main qui lance le jeu
	Droit_POtion_gratuit := true
	clear()          // clear le terminal
	chargement_jeu() // animation de début de jeu (a perfectionner)
	var Max_Item int8 = 10
	var Max_Skill int8 = 2
	PV_Max, Nom, Classe := personnalisation() // fonction qui permet de personnaliser le personnage (Nom d'utilisateur, classe, donc pv max)
	perso1 := InitCaractere(                  // Création du personnage dans la variable perso1 avec la fonction InitCaractère
		Nom,
		Classe,
		1,
		PV_Max,
		100,
		10,                       // Initialisation du nombre d'item max dans l'inventaire du joueur
		2,                        // Initialisation du nombre de sort max du joueur
		make([]Skill, Max_Skill), // Initialisation des sorts du joueur
		make([]Item, Max_Item),   // Initialisation de l'inventaire du joueur
		500,                      // Initialisation de l'argent du joueur
		make([]Equipement, 3),           // Initialisation de l'équipement du joueur
	)
	Add_Inventory_auto(&perso1, &Potion_De_Soin_Nv0, 0) // Ajout d'un item dans l'inventaire du joueur a l'emplacement /10
	Add_Inventory_auto(&perso1, &Epee_En_Bois, 1)       // Ajout d'un item dans l'inventaire du joueur a l'emplacement 1/10
	clear()
	Menu_In_Game(&perso1, Droit_POtion_gratuit)
}

func Menu_In_Game(perso1 *Personnage, Droit_POtion_gratuit bool) Personnage { //affiche le menu Principale du jeu
	var input_Menu string // variable qui permet de naviger dans les sous menus
	fmt.Println("--------------------")
	fmt.Println("1. Information du joueur")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Marchant")
	fmt.Println("4. Forgeron")
	fmt.Println("5. Accèder a vos sorts")
	fmt.Println("6. Accèder a vos équipements")
	fmt.Println("7. Combat-d'Entrainement")
	fmt.Println("8. Quitter le jeu")
	fmt.Println("--------------------")
	fmt.Println("Appuyer sur une touche compris entre 1 et 7 pour naviguer dans le menu")
	fmt.Println("--------------------")
	fmt.Scanln(&input_Menu) //variable qui attend une action de l'utilisateur pour aller dans un sous menu

	switch input_Menu { //attend une action de l'utilisateur

	case "1": //si l'utilisateur appuie sur 1, affiche les informations du personnage
		clear()
		if Droit_POtion_gratuit == true {
			Display_Info_Perso1(perso1, Droit_POtion_gratuit) // affiche les informations du perso1 
			var input_Menu_1 string
			fmt.Println("Appuyer sur une touche pour revenir au menu")
			fmt.Scanln(&input_Menu_1)
			if input_Menu_1 != " " { // retour menu in game 
				clear()
				Menu_In_Game(perso1, Droit_POtion_gratuit)
			}
		} else if Droit_POtion_gratuit == false {
			fmt.Println("Vous avez déja reçu votre potion de vie gratuite")
		}
	case "2": //si l'utilisateur appuie sur 2, affiche l'invantaire
		clear()
		Acces_Inventoy(perso1, Droit_POtion_gratuit)
	case "3": // affiche le shop
		clear()
		Marchant(perso1, Droit_POtion_gratuit)

	case "4": // affiche le forgeron
		clear()
		Forgeron(perso1, Droit_POtion_gratuit)

	case "5": // affiche les sorts du joueur
		clear()
		Display_Info_Sorts(perso1, Droit_POtion_gratuit)

	case "6": 
		//Acces_Equipement(perso1, Droit_POtion_gratuit)
	case "7": //Combat d'entrainement
		Menu_Choise_fight_training(perso1,Droit_POtion_gratuit)
	case "8": // quitter le jeu
		clear()
		fmt.Println("Merci d'avoir joué a notre jeu!!!")
		break

	default: //si l'utilisateur appuie sur une autre touche, affiche une erreur
		clear()
		fmt.Println("Erreur, veuillez appuyer sur une touche compris entre 1 et 4")
		time.Sleep(2 * time.Second)
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	}
	return *perso1
}

func personnalisation() (int16, string, string) { // fonction qui permet de aux joueurx de choisir leur nom d'utilisateur et leur classe
	var Nom string
	var Classe string
	var PV_Max int16
	clear()
	fmt.Println("Character creation in progress...")
	time.Sleep(1 * time.Second)
	clear()
	fmt.Print("Please enter a name: ")
	fmt.Scanln(&Nom)
	if Name_User_Verification(Nom) == true {
		clear()
		fmt.Println("Name,Validé")
		Nom = ConvertToLowerCase(Nom)
		Nom = CapitalizeFirstLetter(Nom)
		time.Sleep(1 * time.Second)
	} else if Name_User_Verification(Nom) == false {
		clear()
		fmt.Println("Name,incorrect")
		time.Sleep(1 * time.Second)
		clear()
		personnalisation()
	}
	clear()
	for {
		fmt.Print("Please choose a category between Human, Elf, and Dwarf: ")
		fmt.Scanln(&Classe)
		switch Classe {
		case "Human":
			// Initialisation des variables en fonction de la classe
			PV_Max = 100
		case "Elf":
			// Initialisation des variables en fonction de la classe
			PV_Max = 80
		case "Dwarf":
			// Initialisation des variables en fonction de la classe
			PV_Max = 120
		default:
			clear()
			fmt.Println("Invalid Category. Please choose again.")
			continue // Continuez à demander la classe tant qu'elle est invalide
		}
		// Si la classe est valide, sortez de la boucle
		return PV_Max, Nom, Classe
	}
}
/*
forgeron a faire
marchnt a finir OK!
mana a faire
affiche sort a faire
combat a faire
equiepement fait 



*/
