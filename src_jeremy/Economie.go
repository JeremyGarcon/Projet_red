package main

func GoldCheck(perso1 *Personnage, Somme int) bool {
	if perso1.Argent <= Somme {
		return false
	}
	return true
}
