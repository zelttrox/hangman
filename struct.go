package main

type Hangman struct {
	WordList     []string // Liste des mots
	ListLen      int      // Taille de la liste

	WordLen      int      // Length of the chosen word
	ToFind       string   // Final word chosen by the program at the beginning. It is the word to find
	
	FoundLetters string   // Lettres trouvées
	Attempts     int      // Number of attempts left
}
