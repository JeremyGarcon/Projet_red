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
	fmt.Println("\n")
	clear()
	fmt.Println("\033[32m Once upon a time, in the fantastic world of Camellium, a prophecy was made.\n" + // intro
		"An individual, coming from the sky, would become a hero and accomplish miracles.\n" +
		"During centuries, the different tribes hoped that they would be the chosen ones, and would become the most powerful state to govern the others.\n \033[0m")
	time.Sleep(10 * time.Second)
	
	fmt.Println("\033[32m The first tribe to stand out from the others was the Human tribe. They weren't the strongest, nor the smartest.\n" +
		"Their strength was coming from the other tribes and their trading skills. They had managed to master the art of negociation to a point that no one could win a trade against \n" +
		"them.\n" +
		"After a few dozens of years, their possessions were so massive that they were the most well rounded nation. \n" +
		" Later, came the dwarves. Their physical strength was over the top. They also mastered the arts of the forge, and the individual battle. \n" +
		" They claimed their place as one of the strongest nations after a bloody war against the Humans, the first ever open conflict to oppose two entire nations.\n" +
		"The Dwarves won thanks to their weapon mastery and the quality of their iron. Humans managed to hold their ground still, but had to accept that they weren't the only ones \n" +
		"at the top anymore.\n" +
		"These two nations came back to peace after realizing that war was meaningless and that the Dwarves had proven themselves\n \033[0m")
    time.Sleep(20 * time.Second)

	fmt.Println("\033[32m This peace didn't remain for long, as the third great nation made their apparition.\n" +
		"The Elves, jealous from the Dwarves status, ambushed a Dwarf's ambassadors group. They made the Dwarves prisonners, and declared war.\n" +
		"The Dwarves didn't back down, confident in their strength, but even if they were outnumbering the Elves and had a much better equipment.\n" +
		"They got defeated after nearly 3 days of war. How ? The Elves are experts when it comes to enchanting spells and archery. Their group organisation was much better as well.\n" +
		"After this war concluded, Elves asked to be one of those superior nations too. None of the Humans, nor the Dwarves were against it.\n \033[0m")
	time.Sleep(20 * time.Second)

	fmt.Println("\033[32m We are now hundreds of years later, and the situation stabilized as it is. Some tribes, like the Orcs or the Dark Elves tried to elevate themselves as well, " +
		"but never won a war, resulting in them being bloodthirsty for those three superior nations, and staying minor tribes. \033[0m")
	time.Sleep(10 * time.Second)

	fmt.Println("\033[32m You can now chose your path. Will you be a Human negociator, who masters nothing in particular ? A Dwarf whose fighting skills are near uncomparable ? Or an Elf \n" +
		"who masters many enchants and knows better than anyone the arts of War Strategy ? \033[0m")
	time.Sleep(10 * time.Second)
}
