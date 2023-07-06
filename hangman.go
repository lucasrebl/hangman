package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Random(file_name string) {

	var tab_word_Vrune []rune //mot caché, devient miroir de to_find version "____"
	var tab_word []string     //tableau de word.txt
	var tab_word_letter []string
	life := 10   //nombre de vie
	var win bool //Pour connaitre si le tour est victorieu
	fmt.Println(life)

	// fonction mot aleatoire
	rand.Seed(time.Now().UnixNano()) // seed time

	fmt.Println("\033[2J") // clear terminal

	file_words, _ := os.Open(file_name)          //Ouverture word.txt
	scanner_word := bufio.NewScanner(file_words) //Flux Word.txt

	input := bufio.NewScanner(os.Stdin) //Flux input

	for scanner_word.Scan() { //boucle ligne par ligne
		tab_word = append(tab_word, scanner_word.Text()) //on ajoute le mot dans tab_word_Vrune
	}
	nbr_words := rand.Intn(len(tab_word)) + 1 //genere le nbr aleatoire
	//fmt.Println(nbr_words)//===================== test nbr random
	//fmt.Println(tab_word)//==========================
	to_found := tab_word[nbr_words] //mot à trouver
	//fmt.Println(to_found)

	for i := 0; i < len(to_found); i++ { //creation miroir de to_find version "_"
		tab_word_Vrune = append(tab_word_Vrune, '_')
	}
	tab_word_Vrune = StartRevealLetters(to_found, tab_word_Vrune)
	to_found_Vrune := String_To_Rune(to_found) //appel de String_To_Rune (convert to_found en tab de rune pour Equal)
	// fmt.Println(to_found_Vrune)

	fmt.Println("Good Luck, you have 10 attempts.")
	for {
		win = false
		fmt.Println(string(tab_word_Vrune)) // affiche le mot a trouvé avec des '_'
		fmt.Println()
		input.Scan()
		mon_flux := input.Text() //corespond a la lettre que le joueur rentre
		tab_word_letter = append(tab_word_letter, mon_flux)
		// for x := 0; x < len(tab_word_letter); x++ {
		// 	if tab_word_letter(x) == []mon_flux {
		// 		fmt.Println(tab_word_letter)
		// 	}
		// }
		fmt.Println("lettre deja utilisé :", tab_word_letter)

		for j := 0; j < len(to_found); j++ { // boucle qui permet d'afficher la lettre rentrée par l'utilisateur
			if mon_flux == (string(to_found[j])) { // condition qui affiche la lettre rentréé autant de fois qu'elle est dans le mot
				win = true
				tab_word_Vrune[j] = (rune(mon_flux[0]))
			}
		}

		if len(mon_flux) > 1 {
			win, tab_word_Vrune, life = IsWord(to_found, mon_flux, tab_word_Vrune, life) // appel de la fonction IsWord
		} else {
			win, tab_word_Vrune, life = IfLetterInTheWord(to_found, mon_flux, tab_word_Vrune, life) // appel de la fonction IfLetterInTheWord
		}
		if !win {
			fmt.Println(life)
			//print hangman
			if Read("hangamn.txt")[9] == byte(13) { // condition qui affiche le pendu si lettre ou mot faux
				fmt.Println("\033[2J") // clear terminal
				fmt.Println(string(Read("hangman.txt")[(9-life)*79 : (9-life)*79+77]))
			} else {
				fmt.Println("\033[2J") // clear terminal
				fmt.Println(string(Read("hangman.txt")[(9-life)*71 : (9-life)*71+70]))
			}
			//fin print hangman
			if life >= 1 {
				fmt.Println("Not present in the word, ", life, " attempts remaining")
			}
		}
		if Equal(to_found_Vrune, tab_word_Vrune) {
			fmt.Println("\033[2J") // clear terminal
			fmt.Println("YYYYYYY       YYYYYYY     OOOOOOOOO     UUUUUUUU     UUUUUUU     WWWWWWWW                           WWWWWWWWIIIIIIIIII	NNNNNNNN        NNNNNNNN ")
			fmt.Println("Y:::::Y       Y:::::Y   OO:::::::::OO   U::::::U     U:::::U     W::::::W                           W::::::WI::::::::I	N:::::::N       N::::::N ")
			fmt.Println("Y:::::Y       Y:::::Y OO:::::::::::::OO U::::::U     U:::::U     W::::::W                           W::::::WI::::::::I	N::::::::N      N::::::N ")
			fmt.Println("Y::::::Y     Y::::::YO:::::::OOO:::::::OUU:::::U     U:::::U      W::::::W                           W::::::WII::::::I	IN:::::::::N     N::::::N ")
			fmt.Println("  Y:::::Y   Y:::::Y  O::::::O   O::::::O U:::::U     U:::::U       W:::::W           WWWWW           W:::::W   I::::I  N::::::::::N    N::::::N ")
			fmt.Println("   Y:::::Y Y:::::Y   O:::::O     O:::::O U:::::D     D:::::U        W:::::W         W:::::W         W:::::W    I::::I  N:::::::::::N   N::::::N ")
			fmt.Println("    Y:::::Y:::::Y    O:::::O     O:::::O U:::::D     D:::::U         W:::::W       W:::::::W       W:::::W     I::::I  N:::::::N::::N  N::::::N ")
			fmt.Println("      Y:::::::Y      O:::::O     O:::::O U:::::D     D:::::U           W:::::W   W:::::W:::::W   W:::::W       I::::I  N::::::N  N::::N:::::::N ")
			fmt.Println("       Y:::::Y       O:::::O     O:::::O U:::::D     D:::::U            W:::::W W:::::W W:::::W W:::::W        I::::I  N::::::N   N:::::::::::N ")
			fmt.Println("       Y:::::Y       O:::::O     O:::::O U:::::D     D:::::U            W:::::W W:::::W W:::::W W:::::W        I::::I  N::::::N   N:::::::::::N ")
			fmt.Println("       Y:::::Y       O::::::O   O::::::O U::::::U   U::::::U              W:::::::::W     W:::::::::W          I::::I  N::::::N     N:::::::::N ")
			fmt.Println("       Y:::::Y       O:::::::OOO:::::::O U:::::::UUU:::::::U               W:::::::W       W:::::::W         II::::::IIN::::::N      N::::::::N ")
			fmt.Println("    YYYY:::::YYYY     OO:::::::::::::OO   UU:::::::::::::UU                 W:::::W         W:::::W          I::::::::IN::::::N       N:::::::N ")
			fmt.Println("    Y:::::::::::Y       OO:::::::::OO       UU:::::::::UU                    W:::W           W:::W           I::::::::IN::::::N        N::::::N ")
			fmt.Println("    YYYYYYYYYYYYY         OOOOOOOOO           UUUUUUUUU                       WWW             WWW            IIIIIIIIIINNNNNNNN         NNNNNNN ")
			fmt.Println("congratulations you win")
			fmt.Println("the word you found is: ", to_found)
			break
		} else if life == 0 {
			fmt.Println("                                                                 lllllll                                                                        ")
			fmt.Println("                                                                 l:::::l                                                                        ")
			fmt.Println("                                                                 l:::::l                                                                        ")
			fmt.Println("                                                                 l:::::l                                                                        ")
			fmt.Println("yyyyyyy           yyyyyyy   ooooooooooo   uuuuuu    uuuuuu        l::::l    ooooooooooo      ooooooooooo       ssssssssss       eeeeeeeeeeee    ")
			fmt.Println(" y:::::y         y:::::y  oo:::::::::::oo u::::u    u::::u        l::::l  oo:::::::::::oo  oo:::::::::::oo   ss::::::::::s    ee::::::::::::ee  ")
			fmt.Println("  y:::::y       y:::::y  o:::::::::::::::ou::::u    u::::u        l::::l o:::::::::::::::oo:::::::::::::::oss:::::::::::::s  e::::::eeeee:::::ee")
			fmt.Println("   y:::::y     y:::::y   o:::::ooooo:::::ou::::u    u::::u        l::::l o:::::ooooo:::::oo:::::ooooo:::::os::::::ssss:::::se::::::e     e:::::e")
			fmt.Println("    y:::::y   y:::::y    o::::o     o::::ou::::u    u::::u        l::::l o::::o     o::::oo::::o     o::::o s:::::s  ssssss e:::::::eeeee::::::e")
			fmt.Println("     y:::::y y:::::y     o::::o     o::::ou::::u    u::::u        l::::l o::::o     o::::oo::::o     o::::o   s::::::s      e:::::::::::::::::e ")
			fmt.Println("      y:::::y:::::y      o::::o     o::::ou::::u    u::::u        l::::l o::::o     o::::oo::::o     o::::o      s::::::s   e::::::eeeeeeeeeee  ")
			fmt.Println("       y:::::::::y       o::::o     o::::ou:::::uuuu:::::u        l::::l o::::o     o::::oo::::o     o::::ossssss   s:::::s e:::::::e           ")
			fmt.Println("        y:::::::y        o:::::ooooo:::::ou:::::::::::::::uu     l::::::lo:::::ooooo:::::oo:::::ooooo:::::os:::::ssss::::::se::::::::e          ")
			fmt.Println("        y:::::y         o:::::::::::::::o u:::::::::::::::u     l::::::lo:::::::::::::::oo:::::::::::::::os::::::::::::::s  e::::::::eeeeeeee   ")
			fmt.Println("       y:::::y           oo:::::::::::oo   uu::::::::uu:::u     l::::::l oo:::::::::::oo  oo:::::::::::oo  s:::::::::::ss    ee:::::::::::::e   ")
			fmt.Println("      y:::::y              ooooooooooo       uuuuuuuu  uuuu     llllllll   ooooooooooo      ooooooooooo     sssssssssss        eeeeeeeeeeeeee   ")
			fmt.Println("     y:::::y                                                                                                                                    ")
			fmt.Println("    y:::::y                                                                                                                                     ")
			fmt.Println("   y:::::y                                                                                                                                      ")
			fmt.Println("  y:::::y                                                                                                                                       ")

			fmt.Println("Oh no you loose don't worry and try again")
			fmt.Println("the word you had to find was: ", to_found)
			break
		}

	}

}
