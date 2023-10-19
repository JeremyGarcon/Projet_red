package main

import (
	"fmt"
	"time"
)

/*
Explication du programme :
Ce programme defini tous les fonctions en lien avec l'invantaire du joueur et de la modification du contenue de cet inbvantaire
*/

func Acces_Inventoy(perso1 *Personnage, Droit_POtion_gratuit bool) { //affiche l'inventaire du joueur
	fmt.Println("Inventory: ")
	fmt.Println("-------------------------")
	for i := 0; i < len(perso1.Inventaire); i++ {
		fmt.Printf("%d. %s \n", i, perso1.Inventaire[i].Name)
	}
	fmt.Println("-------------------------")
	fmt.Println("Appuyer sur une touche pour revenir au menu")
	var input_Menu_1 string
	fmt.Scanln(&input_Menu_1)
	if input_Menu_1 != " " {
		clear()
		Menu_In_Game(perso1, Droit_POtion_gratuit)
	}

}
func (perso1 *Personnage) UpgradeInventory() {
    // Vérifier si le personnage a assez d'argent pour l'amélioration
    if perso1.Argent >= 30 {
        // Créer une nouvelle tranche (slice) avec une capacité augmentée
        newInventory := make([]Item, perso1.Max_Item+10)

        // Copier les éléments de l'ancien inventaire vers le nouveau
        for i := 0; i < int(perso1.Max_Item); i++ {
            if i < len(perso1.Inventaire) {
                newInventory[i] = perso1.Inventaire[i]
            }
        }

        // Mettre à jour la référence de l'inventaire du personnage
        perso1.Inventaire = newInventory

        // Augmenter la capacité maximale de l'inventaire
        perso1.Max_Item += 10

        // Soustraire le coût de l'amélioration de l'argent du personnage
        perso1.Argent -= 30
        clear()
        fmt.Println("Inventaire amélioré. Nouvelle capacité :", perso1.Max_Item)
        time.Sleep(2 * time.Second)
    } else {
        clear()
        fmt.Println("Pas assez d'argent pour améliorer l'inventaire.")
        time.Sleep(2 * time.Second)
    }
}


func Add_Inventory_auto(perso1 *Personnage, Item_add *Item, emplacement int) { // fonction qui ajoute un item dans l'inventaire du joueur au démarrage du jeu

	perso1.Inventaire[emplacement] = *Item_add

}

func AddItemToInventory(perso1 *Personnage, item Item) bool {
    clear()
    fmt.Println("Tu dois choisir où tu veux mettre l'item dans ton inventaire.")
    fmt.Println("Voici ton inventaire :")
    fmt.Println("-------------------------")
    for i := 0; i < len(perso1.Inventaire); i++ {
        fmt.Printf("%d. %s\n", i, perso1.Inventaire[i].Name)
    }
    fmt.Println("-------------------------")

    var emplacement int
    fmt.Print("Choisissez un emplacement (0 - ", perso1.Max_Item-1, "): ")
    _, err := fmt.Scan(&emplacement)
    
    if err != nil {
        fmt.Println("Erreur de saisie :", err)
        return false
    }

    if emplacement >= 0 && emplacement < int(perso1.Max_Item) {
        if perso1.Inventaire[emplacement].Name == "" {
            perso1.Inventaire[emplacement] = item
            fmt.Printf("Item ajouté à l'emplacement %d\n", emplacement)
            time.Sleep(2 * time.Second)
            clear()
            return true
        } else {
            clear()
            fmt.Println("Cet emplacement est déjà occupé.")
            time.Sleep(2 * time.Second)
            clear()
            AddItemToInventory(perso1, item)
            return false
        }
    } else {
        clear()
        fmt.Println("Emplacement invalide.")
        time.Sleep(2 * time.Second)
        clear()
        return false
    }
    
}



