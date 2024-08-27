package main

import (
	"fmt"
	"os"
)

const MAX_TRIES int = 6

type Hangman struct {
	word    string
	guessed string
	tries   int
}

var hangman Hangman

func generate_random_word() string {
	//randInt := rand.IntN(10000) + 1

	dat, err := os.ReadFile("word_list.txt")
	if err != nil {
		fmt.Println("Couldn't open word list")
		os.Exit(1)
	}

	word_list := string(dat[:])

	for char, s := range word_list {
		if char == '\n' || s == '\n' {
			fmt.Println("nice")
		}
	}

	return ""
}

func main() {
	fmt.Println("Welcome to the Hangman game\n\n")

	word := generate_random_word()
	hangman.word = word

}
