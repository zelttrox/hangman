package hangman

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Import du dictionnaire dans un tableau de string
func GetWordList(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		WordList = append(WordList, scanner.Text())
	}
	return WordList, scanner.Err()
}

// Import du dictionnaire dans un tableau de string
func GetHangmanList(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		HangmanPosition = append(HangmanPosition, scanner.Text())
	}
	return HangmanPosition, scanner.Err()
}

// Choix aléatoire du mot à partir de la liste
func GetWord() {
	Word = WordList[rand.Intn(len(WordList))]
	Word = strings.ToUpper(Word)
}

// Scan des arguments choisis dans le terminal
func GetArg() {
	argHardcore := flag.Bool("hard", false, "Hardcore mode")
	argDictionary := flag.String("list", "lol", "Choose dictionary")
	flag.Parse()

	if *argHardcore {
		IsHardcoreMode = true

		fmt.Println("\033[31m", "Hardcore Mode enabled", "\033[0m")
	}

	if *argDictionary == "lol" {
		Dictionary = "text/league.txt"

	} else if *argDictionary == "games" {
		Dictionary = "text/games.txt"
	}
}

// Initialisation des cases du Hangman
func InitBlankspace() {
	for range Word {
		Blankspace = append(Blankspace, " _")
	}
}

// Initialisation des fonctions
func Init() {

	IsRunning = false

	Attempts = 10
	MaxAttempts = 10
	HangmanLen = 7

	GetArg()

	fmt.Println(ToAsciiArt("HANGMAN"))
	fmt.Println("\033[32m", "Good Luck, Have Fun you have 10 attempts.\n\n", "\033[0m")

	GetWordList(Dictionary)
	GetHangmanList("text/hangman.txt")
	GetWord()

	InitBlankspace()

	fmt.Println(Blankspace, Attempts)
}
