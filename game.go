package main

import (
	"fmt"
	"strings"
)

func (h *Hangman) Run() {

	h.Lives = 10
	for {
		fmt.Printf("Word %s Letter: ", strings.Join(h.Blankspace, ""))
		fmt.Scanln(&h.Input)
		output := strings.ToUpper(h.Input)
		fmt.Println(output)

		for _, inputletter := range h.Input {
			guess := false
			for i, wordletter := range h.Word {
				if inputletter == wordletter {
					h.Blankspace[i] = string(inputletter)
					guess = true
				}
			}

			if !guess {
				h.Lives--
			}
		}
		fmt.Println(h.Blankspace, h.Lives)

		if h.Lives <= 0 {
			fmt.Printf("</3, Word: %s - AHHAHAHA YOU LOST!", h.Word)
			break
		}
		if h.Word == strings.Join(h.Blankspace, "") {
			fmt.Printf(" <3 %d, Word: %s - GG YOU WON! \n", h.Lives, h.Word)
			break
		}
	}
}
