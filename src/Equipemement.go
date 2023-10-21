package main

import (
	"fmt"
	"time"
)

func Acces_Equipement(perso1 *Personnage, Droit_POtion_gratuit bool) {
    fmt.Println("Equipment: ")
    fmt.Println("-------------------------")
    fmt.Printf("Tête: %s \n", perso1.Equipement[0])
    fmt.Printf("Torse: %s \n", perso1.Equipement[1])
    fmt.Printf("Jambe: %s \n", perso1.Equipement[2])
    fmt.Println("-------------------------")
    fmt.Println("Appuyer sur une touche pour revenir au menu")
    var input_Menu_1 string
    fmt.Scanln(&input_Menu_1)
    if input_Menu_1 != " " {
        clear()
        Menu_In_Game(perso1, Droit_POtion_gratuit)
    }
}


	// fonction pour ajouter un item a l'inventaire
func AjouterItemAEquipement(equipement *Equipement, item Item, emplacement string) {
    switch emplacement {
    case "tête":
        equipement.Equipement_de_tete = Chapeau_de_Laventurier
    case "torse":
        equipement.Equipement_de_torse = Tuniques_de_Laventurier
    case "jambe":
        equipement.Equipement_de_jambe = Bottes_de_Laventurier
    default:
        fmt.Println("Emplacement d'équipement invalide.")
    }
}



// effet Equipement sur le personnage quand le joueur équipe un équipement



func (perso1 *Personnage)Equipement_Effet() Personnage {
	if perso1.Equipement[1].Equipement_de_torse.Name == "Tuniques de l'aventurier" {
	perso1.PV_Max += 10
	perso1.PV_Actuels += 10
	time.Sleep(2 * time.Second)
	clear()
	fmt.Println("Vous avez équipé une Tuniques de l'aventurier")
	return *perso1
	}
	if perso1.Equipement[1].Equipement_de_torse.Name != "Tuniques de l'aventurier" {
		perso1.PV_Max += 10
		perso1.PV_Actuels += 10
		return *perso1
		}
	if perso1.Equipement[2].Equipement_de_jambe.Name == "Bottes de l'aventurier" {
		perso1.PV_Max += 10
		perso1.PV_Actuels += 10
	}
	return *perso1
}
// Annulation des effets de l'équipement a utiliser a chaque fois que l'on change d'équipement 
func (perso1 *Personnage)Equipement_No_Effet() Personnage {
if perso1.Equipement[0].Equipement_de_tete.Name != "Chapeau de l'aventurier" {
	perso1.PV_Max -= 10
	perso1.PV_Actuels -= 10
	return *perso1
	}
if perso1.Equipement[1].Equipement_de_torse.Name != "Tuniques de l'aventurier" {
	perso1.PV_Max -= 10
	perso1.PV_Actuels -= 10
	return *perso1
	}
if perso1.Equipement[2].Equipement_de_jambe.Name != "Bottes de l'aventurier" {
	perso1.PV_Max -= 10
	perso1.PV_Actuels -= 10
	return *perso1
	}
	return *perso1
} 

