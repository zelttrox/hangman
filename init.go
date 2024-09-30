package main

import (
	"fmt"
	"strings"

	"math/rand"
)

func (h *HangMan) InitWorldList() {
	h.WordList = []string{"bonjour", "manger", "salut"}

	h.Word = h.WordList[rand.Intn(len(h.WordList))]

}
func (h *HangMan) InitBlankspace() {
	for range h.Word {
		h.blankspace = append(h.blankspace, " _")
	}

}

func (h *HangMan) Hangman() {

	lives := 7
	for {
		fmt.Printf("Word %s Letter: ", strings.Join(h.blankspace, ""))

		fmt.Scanln(&h.input)
		fmt.Println(h.input)

		for _, inputletter := range h.input {
			guess := false
			for i, wordletter := range h.Word {
				if inputletter == wordletter {
					h.blankspace[i] = string(inputletter)
					guess = true
				}
			}

			if !guess {
				lives--
			}
		}
		fmt.Println(h.blankspace, lives)

		if lives <= 0 {
			fmt.Printf("</3, Word: %s - AHHAHAHA YOU LOST!", h.Word)
			break
		}
		if h.Word == strings.Join(h.blankspace, "") {
			fmt.Printf(" <3 %d, Word: %s - GG YOU WON! \n", lives, h.Word)
			break
		}
	}
}
