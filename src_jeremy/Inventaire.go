package main

/* Explication_Du_Programme :
Ce programme crée les structure de l'inventaire et des items
et donne accès a l'inventaire du joueur


type Personnage_intventory struct { // Structure de l'inventaire
	Inventaire []Item
}

type Item struct { //structure des items
	Name        string //SSID de l'item
	id          int    // numéro de l'item
	Value       int    // Valeur de soin ou d'attaque ex: potion = 100 donne 100 pv et a contrario épée = -50 enlève 50 pv a l'adversarie
	description string // description de l'item
}

func Create_Inventory() { //création de l'inventaire du joueur

	joueur := Personnage_intventory{

		Inventaire: []Item{Epee_En_Bois, Potion_De_Soin_Nv0, Potion_De_Poison_Nv0},
	}
	return joueur
}

/*
func accessInventory(*ItemPointer, *PersonnageInventoryPointer, *PersonnagePointer) { // accès a l'inventaire du joueur
	fmt.Println("joueur.Inventaire:")
	for _, item := range joueur.Inventaire {
		fmt.Println(item.Name)
	}
}
*/
