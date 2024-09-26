package main

type HangManData struct {
	Word             string // Word composed of '_', ex: H_ll_
	WordLen          int // Length of the chosen word
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored}
}
