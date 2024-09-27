package main

import (
	"fmt"
	"math/rand"
)

// Import du dictionnaire dans un tableau de string
func (h *Hangman) GetWordList() {
	fmt.Println("~ Importing Words list")
	h.WordList = append(h.WordList, "Hello", "Timing", "Code", "Yes")
}

// Choix d'un mot au hasard à partir du tableau de mots
func (h *Hangman) RandomWord(wordList []string) {
	fmt.Println("~ Choosing a Random word")
	h.ListLen = len(wordList)
	wordChosen := rand.Intn(h.ListLen)
	h.ToFind = wordList[wordChosen]
}

// Initialisation du Hangman
func (h *Hangman) Init() {
	h.GetWordList()
	h.RandomWord(h.WordList)
	fmt.Print("Word: ")
	fmt.Println(h.ToFind) // DEBUG
}
