package main

import (
	"fmt"
	"strings"
)

func (h *Hangman) Run() {

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
			fmt.Printf("</3, Word: %s - AHHAHAHA YOU LOST!", h.Word)
			break
		}
		if h.Word == strings.Join(h.Blankspace, "") {
			fmt.Printf(" <3 %d, Word: %s - GG YOU WON! \n", h.Attempts, h.Word)
			break
		}
	}
}
