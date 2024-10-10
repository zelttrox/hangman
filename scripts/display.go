package hangman

import (
	"fmt"
)

// Progression de l'affichage du pendu
func ProcessHangman() {
	SpaceOut(1)
	for i := HangmanProgress; i < (HangmanLen + HangmanProgress); i++ {
		fmt.Println(HangmanPosition[i])
	}
	HangmanProgress += HangmanLen
}

// Espacement du Terminal
func SpaceOut(lines int) {
	for range lines {
		fmt.Print("\n")
	}
}
