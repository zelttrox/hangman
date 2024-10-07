package hangman

var WordList []string   // Dictionnaire des mots
var Word string         // Mot choisi au hasard
var Blankspace []string // Nombre de cases du Hangman

var HangmanPosition []string // Liste de positions du Hangmans
var HangmanProgress int      // Progression du Hangman
var LinesToAdd int           // Lignes du Hangman à afficher
var HangmanLen int           // Taille d'une frame du Hangman

var Attempts int    // Nombre d'essais restants
var MaxAttempts int // Nombre total d'essais

var Input string   // Lettre choisie par le joueur
var IsRunning bool // Est vrai si le jeu est lancé
