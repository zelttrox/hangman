package main

type Hangman struct {
	Word       string   // Word composed of '_', ex: H_ll_
	Attempts   int      // Nombre d'essais restants
	Input      string   // Lettre choisie par le joueur
	Blankspace []string // Nombre de cases du Hangman
	WordList   []string
}
