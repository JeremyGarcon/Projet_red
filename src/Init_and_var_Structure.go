package main


// init les Items  pour le joueur:
var Epee_En_Bois = Item{ // item de base

	Name:        "Eppée en bois",
	Value:       5, // valeur de l'item sur le pv du joueur ou de l'adversaire
	description: "Une eppée en bois, pas très résistant mais aide a se défendre",
}
var Potion_De_Soin_Nv0 = Item{ // item de base

	Name:        "Potion de soin Nv0",
	Value:       7, // valeur de l'item sur le pv du joueur ou de l'adversaire
	description: "Une potion de soin Nv0, ne regénère que très peu de pv",
}
var Potion_De_Poison_Nv0 = Item{ // item de base

	Name:        "Potion de Poison Nv0",
	Value:       -7, // valeur de l'item sur le pv du joueur ou de l'adversaire
	description: "Une potion de Poison Nv0, ne fait que très peu de dégat aux monstres",
}
var Livre_de_sort = Item{ // item de base

	Name:        "Livre de sort",
	Value:       0, // valeur de l'item sur le pv du joueur ou de l'adversaire
	description: "Un livre de sort, permet d'apprendre de nouveaux sorts ex: Boule de feu",
}
//Item pour la structure Equipement
var Tuniques_de_Laventurier = Item{ // item de base
	Name: 	  "Tuniques de l'aventurier",
	Value: 	  0,
	description: "Tuniques de l'aventurier, permet de se protéger un peu plus",
}
var Bottes_de_Laventurier = Item{ // item de base
	Name:        "Bottes de l'aventurier",
	Value:       0, // valeur de l'item sur le pv du joueur ou de l'adversaire
	description: "Bottes de l'aventurier, permet de se protéger un peu plus",
}
var Chapeau_de_Laventurier = Item{ // item de base
	Name:        "Chapeau de l'aventurier",
	Value:       0, // valeur de l'item sur le pv du joueur ou de l'adversaire
	description: "Chapeau de l'aventurier, permet de se protéger un peu plus",
}

// item de craft
var Fourrure_de_loup = Item{ 
	Name: 	  "Fourrure de loup",
	Value: 	  0,
	description: "Une fourrure de loup, peut être vendu au marchand",

}
var Peau_De_Troll = Item{
	Name: 	  "Peau de troll",
	Value: 	  0,

	description: "Une peau de troll, peut être vendu au marchand",
}
var Plume_De_Corbeau = Item{
	Name: 	  "Peau de corbeau",
	Value: 	  0,
	description: "Une peau de corbeau, peut être vendu au marchand",
}

// sort
var Boule_de_feu = Skill{ // item de base

	Name:        "Boule de feu",
	Value:       -15,
	description: "premier sort appris par le joueur",
}

var Coup_de_poing = Skill{ // item de base

	Name:        "Coupe de poing",
	Value:       -5,
	description: "sort par déaut du joueur",
}
//personnage 

// ini le personnage
func InitCaractere(Nom string, Classe string, Niveau int8, PV_Max int16, PV_Actuels int16, Max_Item int8, Max_Skill int8, Skill []Skill, Inventaire []Item, Argent int, Equipement []Equipement) Personnage {
	return Personnage{
		Nom:        Nom,
		Classe:     Classe,
		Niveau:     Niveau,
		PV_Max:     PV_Max,
		PV_Actuels: PV_Max / 2,
		Max_Item:   Max_Item,
		Max_Skill:  Max_Skill,
		Skill:      Skill,
		Inventaire: Inventaire,
		Argent:     Argent,
		Equipement: Equipement,
	}
}
// fonction pour initialiser les Monstre
func Init_LaChose (Nom string, PV_Max int16, PV_Actuels int16, Attaque int16) Monstre {
	return Monstre{
		Nom:        Nom,
		PV_Max:     PV_Max,
		PV_Actuels: PV_Max,
		Attaque:    Attaque,
	}
}
func Init_Goblin (Nom string, PV_Max int16, PV_Actuels int16, Attaque int16) Monstre {
	return Monstre{
		Nom:        Nom,
		PV_Max:     PV_Max,
		PV_Actuels: PV_Max,
		Attaque:    Attaque,
	}
}
// fonction pour initialiser les Equipements avec des items de base


