package hangman

import (
	"fmt"
)

// HangmanLen = 7
func ProcessHangman() {
	for i := HangmanProgress; i < (HangmanLen + HangmanProgress); i++ {
		fmt.Println(HangmanPosition[i])
	}
	HangmanProgress += HangmanLen
}
