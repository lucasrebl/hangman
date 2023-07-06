package hangman

import (
	"math/rand"
	"time"
)

func IfLetterInTheWord(word string, letter string, tab_rune []rune, life int) (bool, []rune, int) {
	for j := 0; j < len(word); j++ {
		if letter == (string(word[j])) { //vérifi si lettre est dans le mot
			tab_rune[j] = (rune(letter[0])) //vérifi si lettre est dans tab_rune
			return true, tab_rune, life
		}
	}
	return false, tab_rune, life - 1
}
func IsWord(word string, wordentry string, tab_found []rune, life int) (bool, []rune, int) {
	if word == wordentry { //vérifi si mot entrée par utilisateur = mot à trouvé
		tab_found = []rune(word)
		return true, tab_found, life
	}
	return false, tab_found, life - 2
}

func StartRevealLetters(word string, tab_rune []rune) []rune {

	rand.Seed(time.Now().UnixNano())
	nblettrealeatoire := rand.Intn((len(word)-1)/2) + 1 //nb d'indice révéler dans le mot
	for i := 0; i < nblettrealeatoire; i++ {
		randomindex := rand.Intn(len(word)-1) + 1       //choisi lettres aléatoirs en fonction du nb d'indice
		tab_rune[randomindex] = rune(word[randomindex]) //ajoute indice dans tab_rune et compare avec le mot
	}
	return tab_rune
}
