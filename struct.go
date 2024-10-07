package main

type Hangman struct {
	WordList   []string // Dictionnaire des mots
	Word       string   // Mot choisi au hasard
	Blankspace []string // Nombre de cases du Hangman

	HangmanPosition []string // Liste de positions du Hangmans
	HangmanProgress int      // Progression du Hangman
	LinesToAdd      int      // Lignes du Hangman à afficher
	HangmanLen      int      // Taille d'une frame du Hangman

	Attempts    int // Nombre d'essais restants
	MaxAttempts int // Nombre total d'essais

	Input     string // Lettre choisie par le joueur
	IsRunning bool   // Est vrai si le jeu est lancé
}
