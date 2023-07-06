package hangman

import (
	"fmt"
	"os"
)

func Read(fichier string) []byte {
	file, err := os.Open(fichier)
	if err != nil {
		fmt.Println(err)
	}
	arr := make([]byte, 1200)
	file.Read(arr)
	return arr
}
