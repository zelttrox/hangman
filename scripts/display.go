package hangman

import (
	"fmt"
)

// HangmanLen = 7
func (h *Hangman) ProcessHangman() {
	for i := h.HangmanProgress; i < (h.HangmanLen + h.HangmanProgress); i++ {
		fmt.Println(h.HangmanPosition[i])
	}
	h.HangmanProgress += h.HangmanLen
}
