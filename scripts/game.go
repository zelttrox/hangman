package hangman

import (
	"fmt"
	"strings"
)

func CheckAttempted(input string) {
	for i := 0; i < len(AttemptedLetter); i++ {
		if input == AttemptedLetter[i] {
			LetterTried = true
		} else {
			LetterTried = false
		}
	}
}

func Run() {

	IsRunning = true

	for {
		fmt.Scanln(&Input)
		Input = strings.ToUpper(Input)
		fmt.Println(Input)

		for _, inputletter := range Input {
			guess := false
			for i, wordletter := range Word {
				if inputletter == wordletter {
					Blankspace[i] = string(inputletter)
					guess = true
				}
			}

			if !guess {
				CheckAttempted(Input)
				if LetterTried {
					fmt.Println("You've already guessed that letter. Try again!")
				} else if !LetterTried {
					Attempts--
					ProcessHangman()
					AttemptedLetter = append(AttemptedLetter, Input)
				}
			}
		}
		fmt.Println(Blankspace, Attempts)
		fmt.Println("Attempted Letter(s): ", AttemptedLetter)

		if Attempts <= 0 {
			fmt.Println(ToAsciiArt(Word))
			fmt.Print("</3, - AHHAHAHA YOU LOST!", Word)
			break
		}
		if Word == strings.Join(Blankspace, "") {
			fmt.Println(ToAsciiArt(Word))
			fmt.Printf(" <3 %d, - GG YOU WON! \n", Attempts)
			break
		}
	}
}
