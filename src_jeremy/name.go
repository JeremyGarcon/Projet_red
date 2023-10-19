package main

import "strings"

func Name_User_Verification(s string) bool {
    for _, char := range s {
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
            continue // Caractère alphabétique, continuez la boucle
        } else {
            return false // Caractère spécial trouvé, renvoyez false immédiatement
        }
    }
    return true // Aucun caractère spécial trouvé, renvoyez true
}

func ConvertToLowerCase(input string) string {
	return strings.ToLower(input)
}

func CapitalizeFirstLetter(input string) string {
	return strings.Title(input)
}
