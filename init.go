package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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
		h.List = append(h.List, scanner.Text())
	}
	return h.List, scanner.Err()
}

func (h *Hangman) RandomWord() {
	h.WordList = h.List
	h.Word = h.WordList[rand.Intn(len(h.WordList))]
}

func (h *Hangman) InitBlankspace() {
	for range h.Word {
		h.Blankspace = append(h.Blankspace, " _")
	}
}

func (h *Hangman) Init() {
	h.GetWordList("words.txt")
	h.RandomWord()
	h.InitBlankspace()
}
