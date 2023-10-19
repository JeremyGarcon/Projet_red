package main

import "fmt"

func Dead(perso1 *Personnage) { // fonction qui vérifie si le joueur est mort (a utiliser lors des combats)
	if perso1.PV_Actuels == 0 {
		clear()
		fmt.Println("Vous êtes mort")
	}
}
