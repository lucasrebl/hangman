package hangman

import (
	"fmt"
)

func Level() string {
	var level string

	fmt.Println("Enter the level of difficulty: ")
	fmt.Println("choose the difficulty between: easy, medium and hard")
	fmt.Scan(&level)

	if level == "easy" { // condition qui choisit le niveau de difficulté
		//Easy(level) // si la condition est valide ouverture du fichier word1.txt
		return "word1.txt"
	} else if level == "medium" {
		//Medium(level) // si la condition est valide ouverture du fichier word2.txt
		return "word2.txt"
	} else if level == "hard" {
		//Hard(level) // si la condition est valide ouverture du fichier word3.txt
		return "word3.txt"
	} else {
		fmt.Println("Wrong input") // sinon refusé l'acces //retourne un invalide argument
		return ("")
	}

}
