package hangman

import (
	"fmt"
	"strings"
)

// Vérifie si la lettre entrée a déjà été essayée
func CheckAttempted(input string) {
	for i := 0; i < len(AttemptedLetter); i++ {
		if input == AttemptedLetter[i] {
			LetterTried = true
		}
	}
}

// Fonction principale du jeu
func Run() {

	IsRunning = true

	for {
		fmt.Scanln(&Input)
		Input = strings.ToUpper(Input)

		// Vérifie si l'input est un mot ou une lettre
		if len(Input) > 1 {
			if Input == Word {
				Win()
			} else if Input != Word {
				Attempts -= 2
				HangmanProgress += HangmanLen
				ProcessHangman()
				SpaceOut(1)
				fmt.Println("\033[33m", "The word is not", Input, "Try again!", "\033[0m")
				fmt.Println(" You lost 2 attemps.", "\n", "")
			}
		}

		// Ajout des lettres devinées dans le mot
		for _, inputletter := range Input {
			guess := false
			for i, wordletter := range Word {
				if inputletter == wordletter {
					Blankspace[i] = string(inputletter)
					guess = true
				}
			}

			// Vérifie si la lettre entrée a déjà été essayée ou si elle est fausse
			if !guess && len(Input) == 1 {
				CheckAttempted(Input)
				if LetterTried {
					fmt.Println("You've already guessed that letter. Try again!")
					LetterTried = false
				} else if !LetterTried {
					Attempts--
					ProcessHangman()
					AttemptedLetter = append(AttemptedLetter, Input)
				}
			}
		}
		// Affiche les lettres déjà essayées
		fmt.Println(Blankspace, Attempts)
		fmt.Println("Attempted Letter(s): ", "\033[36m", AttemptedLetter, "\033[0m")
		SpaceOut(4)

		// Condition de perte
		if Attempts <= 0 {
			Lose()
			break
		}
		// Condition de victoire
		if Word == strings.Join(Blankspace, "") {
			Win()
			break
		}
	}
}

// Fonction de perte
func Lose() {
	fmt.Println(ToAsciiArt(Word))
	fmt.Print("</3, - AHHAHAHA YOU LOST!", Word)
}

// Fonction de victoire
func Win() {
	fmt.Println(ToAsciiArt(Word))
	fmt.Printf(" <3 %d, - GG YOU WON! \n", Attempts)
}
