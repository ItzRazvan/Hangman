package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const MAX_TRIES int = 6

type Hangman struct {
	word    string
	guessed string
	tries   int
}

var hangman Hangman

func generate_random_word() string {
	rand.Seed(time.Now().UnixNano())

	dat, err := os.ReadFile("word_list.txt")
	if err != nil {
		fmt.Println("Couldn't open word list")
		os.Exit(1)
	}

	word_list := strings.Split(string(dat), "\n")

	if len(word_list) == 0 {
		fmt.Println("Word list is empty")
		os.Exit(1)
	}

	random_index := rand.Intn(len(word_list))
	word := word_list[random_index]

	return word
}

func guessed_init(word string) string {
	guessed := ""
	for i := 0; i < len(word); i++ {
		guessed += "_"
	}
	return guessed
}

func main() {
	fmt.Println("Welcome to the Hangman game\n\n")

	word := generate_random_word()
	guessed := guessed_init(word)

	fmt.Println(word, guessed)

}
