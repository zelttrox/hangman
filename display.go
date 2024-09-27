package main

import "fmt"

// Affichage des cases du Hangman en fonction de la taille du mot
func PrintBlank(word string) {
	fmt.Println("~ Affichage des cases du mot")
	for range word {
		fmt.Print("_  ")
	}
}

func PrintLetter(word string) {
	
}

// Fonction d'affichage du Hangman
func Display() {
	PrintBlank(h.ToFind)
}
