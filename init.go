package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Import du dictionnaire dans un tableau de string
func (h *Hangman) GetWordList(path string) ([]string, error) {
	fmt.Println("~ Importing Words list")

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		h.WordList = append(h.WordList, scanner.Text())
	}
	return h.WordList, scanner.Err()
}

// Choix aléatoire du mot à partir de la liste
func (h *Hangman) GetWord() {
	h.Word = h.WordList[rand.Intn(len(h.WordList))]
	h.Word = strings.ToUpper(h.Word)
}

// Initialisation des cases du Hangman
func (h *Hangman) InitBlankspace() {
	for range h.Word {
		h.Blankspace = append(h.Blankspace, " _")
	}
}

// Initialisation des fonctions
func (h *Hangman) Init() {
	h.GetWordList("words.txt")
	h.GetWord()
	h.InitBlankspace()
}
