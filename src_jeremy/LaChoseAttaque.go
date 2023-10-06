package main

import (
    "fmt"
    "time"
)

func (perso1 *Personnage) Attaque_Chose(LaChose_training *Monstre) { // Paterne de l'ennemie
    var Attaque_Chose int = 0
    fmt.Println("Le troll attaque le perso1 !")
    if Attaque_Chose < 4 {
        perso1.Attaque_Simple(LaChose_training)
        Attaque_Chose++
    } else if Attaque_Chose == 4 {
        perso1.Attaque_Speciale(LaChose_training)
        Attaque_Chose = 0
    }
}

func (perso1 *Personnage) Attaque_Simple(LaChose_training *Monstre) {
    fmt.Println("Le troll attaque le perso1 !avec attaque simple")
    attackAnimation()
	clear()
	perso1.PV_Actuels -= LaChose_training.Attaque // Attaque simple
    fmt.Println("Il vous reste %d sur %d",perso1.PV_Actuels, perso1.PV_Max)
}

func (perso1 *Personnage) Attaque_Speciale(LaChose_training *Monstre) {
    fmt.Println("Le troll attaque le perso1 !avec attaque spéciale")
    attackAnimation()
	clear()
	perso1.PV_Actuels -= LaChose_training.Attaque * 2 // Attaque spéciale
    fmt.Println("Il vous reste %d sur %d",perso1.PV_Actuels, perso1.PV_Max)
}

func attackAnimation() {
    for i := 0; i < 5; i++ {
        clear()
        if i%2 == 0 {
            fmt.Println("     \\|/")
        } else {
            fmt.Println("      |")
        }
        fmt.Println("     / \\")
        time.Sleep(500 * time.Millisecond)
    }
    clear()
    fmt.Println("Le troll attaque l'humain !")
}

