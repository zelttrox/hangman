package main

type Hangman struct {
	Word       string // Word composed of '_', ex: H_ll_
	WordList   []string
	Lives      int
	Input      string
	Blankspace []string
	List       []string

	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}
