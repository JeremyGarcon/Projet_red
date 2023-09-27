package main

import (
    "fmt"
)

// Structure représentant l'inventaire du joueur
type Inventory struct {
    Items map[string]int
}

// Fonction pour ajouter un article à l'inventaire
func (inv *Inventory) AddItem(item Item) {
    if inv.Items == nil {
        inv.Items = make(map[string]int)
    }
    inv.Items[item.Name]++
}

// Fonction pour retirer un article de l'inventaire
func (inv *Inventory) RemoveItem(item Item) {
    if count, ok := inv.Items[item.Name]; ok && count > 0 {
        inv.Items[item.Name]--
    }
}

func main() {
    // Créer un inventaire vide pour le joueur
    playerInventory := &Inventory{}

    // Créer un marchand
    merchant := []Item{
        {"Potion de vie", 0},
    }

    // Afficher le menu du marchand
    fmt.Println("Bienvenue chez le marchand ! Voici ce qu'il vend :")
    for i, item := range merchant {
        fmt.Printf("%d. %s (Prix: %d pièces d'or)\n", i+1, item.Name, item.Price)
    }

    // Demander au joueur de faire un choix
    var choice int
    fmt.Print("Veuillez choisir un article à acheter (0 pour quitter) : ")
    fmt.Scan(&choice)

    // Vérifier si le choix est valide
    if choice < 0 || choice > len(merchant) {
        fmt.Println("Choix invalide.")
        return
    }

    // Acheter l'article choisi
    if choice == 0 {
        fmt.Println("Merci d'être passé chez le marchand.")
    } else {
        itemToBuy := merchant[choice-1]
        if itemToBuy.Price > 0 {
            fmt.Printf("Vous avez acheté %s.\n", itemToBuy.Name)
            playerInventory.AddItem(itemToBuy)
        } else {
            fmt.Printf("Vous avez reçu gratuitement %s.\n", itemToBuy.Name)
            playerInventory.AddItem(itemToBuy)
        }
    }

    // Afficher l'inventaire mis à jour du joueur
    fmt.Println("Votre inventaire :")
    for item, count := range playerInventory.Items {
        fmt.Printf("%s: %d\n", item, count)
    }
}

