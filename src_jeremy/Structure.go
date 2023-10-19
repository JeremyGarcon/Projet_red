package main

//Structure Item :

type Item struct { //structure des items
	Name        string //SSID de l'item
	Value       int    // Valeur de soin ou d'attaque ex: potion = 100 donne 100 pv et a contrario épée = -50 enlève 50 pv a l'adversarie
	description string // description de l'item
}
// Structure pour les Sort 

type Skill struct {
	Name        string //SSID de l'item
	Value       int    // Valeur de soin ou d'attaque ex: potion = 100 donne 100 pv et a contrario épée = -50 enlève 50 pv a l'adversarie
	description string // description de l'item
}

// Structure personnage :

type Personnage struct { // Structure pour le personnage
	Nom        string
	Classe     string
	Niveau     int8
	PV_Max     int16
	PV_Actuels int16
	Max_Item   int8
	Max_Skill  int8
	Skill      []Skill
	Inventaire []Item
	Argent     int
	Equipement []Equipement
}

// Structure pour les items
type Inventaire struct { // Structure pour l'inventaire
	Items []Item
}
// Structure pour les équipement 
type Equipement struct { // Structure pour l'équipement
	Equipement_de_tete Item
	Equipement_de_torse Item
	Equipement_de_jambe Item
}

// Structure pour les Bots
type Monstre struct { // Structure pour les items
	Nom        string //nom du bot
	PV_Max     int16  //PV max du bot
	PV_Actuels int16 //PV actuel du bot
	Attaque    int16 //attaque du bot

}
