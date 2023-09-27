package main

import (
	"fmt"
	"time"
)

func takePot(personnage *Personnage) { // Fonction pour utiliser une potion dans l'inventaire

	var potIndex int = -1 // Recherche d'une potion dans l'inventaire
	for i, item := range personnage.Inventaire {
		if item.Name == "Potion" {
			potIndex = i
			break
		}
	}

	if potIndex != -1 {

		personnage.Inventaire = append(personnage.Inventaire[:potIndex], personnage.Inventaire[potIndex+1:]...) // Utilisation de la potion
		personnage.PV_Actuels += 50                                                                             // Ajout de 50 points de vie actuels

		if personnage.PV_Actuels > personnage.PV_Max { // Vérification pour ne pas dépasser le maximum de pv
			personnage.PV_Actuels = personnage.PV_Max
		}
		fmt.Println("Vous avez utilisé une potion et récupéré 50 points de vie.")
	} else {
		fmt.Println("Aucune potion trouvée dans l'inventaire.")
	}

	fmt.Printf("Points de vie actuels : %d / %d\n", personnage.PV_Actuels, personnage.PV_Max)
}


// Fonction pour infliger des dégâts de poison au personnage
func poisonPot(personnage *Personnage) {
	fmt.Println("Vous avez utilisé une Potion de poison.")
	for i := 0; i < 3; i++ {
		personnage.PV_Actuels -= 10
		if personnage.PV_Actuels < 0 {
			personnage.PV_Actuels = 0
		}
		fmt.Printf("Points de vie actuels : %d / %d\n", personnage.PV_Actuels, personnage.PV_Max)
		time.Sleep(1 * time.Second)
	}
}

// Fonction pour acheter une potion auprès du marchand
func acheterPotion(personnage *Personnage, marchand *Personnage) {
	fmt.Println("Vous avez acheté une Potion de poison auprès du marchand.")
	Potion := Item{"Potion de poison"}
	personnage.Inventaire = append(personnage.Inventaire)
}
