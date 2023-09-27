package main

func main() {
	// playMusic() en cours
	var joueurAsPersonnage bool = false //variable qui permet de savoir si le joueur a un personnage
	var Pointeur_joueurAsPersonnage *bool = &joueurAsPersonnage
	chargement_jeu()                        //charge le jeu
	Menu_Start(Pointeur_joueurAsPersonnage) //affiche le menu de depart
}

/*
func playMusic() {
	// Utilisez "wmplayer.exe" pour Windows avec le fichier MP3 à la racine
	cmd := exec.Command("wmplayer.exe", "musique.mp3")

	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}
*/
