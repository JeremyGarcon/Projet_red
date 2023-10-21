package main

import (
    "fmt"
    "time"
)

func chargement_jeu() {
    anime()
    const totalPourcentage = 100
    const delayBetweenPourcentage = 1500 * time.Millisecond

    clear()
    fmt.Println("------ 0 % ------")

    for pourcentage := 25; pourcentage <= totalPourcentage; pourcentage += 25 {
        time.Sleep(delayBetweenPourcentage)
        clear()
        fmt.Printf("------ %d %% ------\n", pourcentage)
    }

    clear()
    fmt.Println("\033[32mOnce upon a time, in the fantastic world of Camellium, a prophecy was made.\n" +
        "An individual, coming from the sky, would become a hero and accomplish miracles.\n" +
        "During centuries, the different tribes hoped that they would be the chosen ones, and would become the most powerful state to govern the others.\n\033[0m")
    time.Sleep(1 * time.Second)
    clear()

    // Le reste de votre animation
}

func anime() {
    animation := []string{"|", "/", "-", "\\"}

    for i := 0; i < 5; i++ {
        for _, frame := range animation {
            fmt.Printf("\r%s", frame)
            time.Sleep(100 * time.Millisecond)
        }
    }
    fmt.Println("\n")
}
