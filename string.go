package hangman

func String_To_Rune(to_found string) []rune {
	var result []rune
	for i := 0; i < len(to_found); i++ {
		result = append(result, (rune(to_found[i])))
	}
	return result
}
