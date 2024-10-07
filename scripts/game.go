package hangman

import (
	"fmt"
	"strings"
)

func Run() {

	IsRunning = true

	for {
		fmt.Printf("Word %s Letter: ", strings.Join(Blankspace, ""))
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
				Attempts--
				ProcessHangman()
			}
		}
		fmt.Println(Blankspace, Attempts)

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
