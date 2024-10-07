package hangman

import (
	"fmt"
	"strings"
)

func Run() {

	var h Hangman

	h.IsRunning = true

	for {
		fmt.Printf("Word %s Letter: ", strings.Join(h.Blankspace, ""))
		fmt.Scanln(&h.Input)
		h.Input = strings.ToUpper(h.Input)
		fmt.Println(h.Input)

		for _, inputletter := range h.Input {
			guess := false
			for i, wordletter := range h.Word {
				if inputletter == wordletter {
					h.Blankspace[i] = string(inputletter)
					guess = true
				}
			}

			if !guess {
				h.Attempts--
				h.ProcessHangman()
			}
		}
		fmt.Println(h.Blankspace, h.Attempts)

		if h.Attempts <= 0 {
			fmt.Println(h.ToAsciiArt(h.Word))
			fmt.Print("</3, - AHHAHAHA YOU LOST!", h.Word)
			break
		}
		if h.Word == strings.Join(h.Blankspace, "") {
			fmt.Println(h.ToAsciiArt(h.Word))
			fmt.Printf(" <3 %d, - GG YOU WON! \n", h.Attempts)
			break
		}
	}
}
