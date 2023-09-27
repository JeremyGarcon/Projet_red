package main

import (
	"fmt"
	"time"
)

func chargement_jeu() { // fonction qui affiche un chargement de début de jeu
	anime()
	for pourcentage := 0; pourcentage <= 100; pourcentage = pourcentage + 25 {
		clear()
		fmt.Println("------", pourcentage, "%", "------")
		time.Sleep(1 * time.Second)
		clear()
	}
}
func anime() { // fonction qui affiche un chargement
	animation := []string{"|", "/", "-", "\\"}

	for i := 0; i < 5; i++ {
		for _, frame := range animation {
			fmt.Printf("\r%s", frame)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
