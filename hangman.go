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

func hangman_init(word string, guessed string) {
	hangman.word = word
	hangman.guessed = guessed
	hangman.tries = 0
}

func print_hangman() {
	switch hangman.tries {
	case 0:
		fmt.Println("------------\n|          |\n|\n|\n|\n|\n|\n|\n|\n|\n|\n|\n|")
	case 1:
		fmt.Println("------------\n|          |\n|         ---\n|        |. .|\n|         ---\n|\n|\n|\n|\n|\n|\n|")
	case 2:
		fmt.Println("------------\n|          |\n|         ---\n|        |   |\n|         ---\n|          |\n|          |\n|          |\n|\n|\n|\n|")
	case 3:
		fmt.Println("------------\n|          |\n|         ---\n|        |   |\n|         ---\n|         /|\n|          |\n|          |\n|\n|\n|\n|")
	case 4:
		fmt.Println("------------\n|          |\n|         ---\n|        |   |\n|         ---\n|         /|\\ \n|          |\n|          |\n|\n|\n|\n|")
	case 5:
		fmt.Println("------------\n|          |\n|         ---\n|        |   |\n|         ---\n|         /|\\ \n|          |\n|          |\n|         /\n|\n|\n|")
	default:
	}
}

func print_elements() {
	fmt.Println("Word: ", hangman.guessed)
	print_hangman()
}

func print_lose_screen() {
	fmt.Println("   YOU LOST\n")
	fmt.Println("You guessed: ", hangman.guessed)
	fmt.Println("The correct word: ", hangman.word)
	fmt.Println("------------\n|          |\n|         ---\n|        |. .|\n|         ---\n|         /|\\ \n|          |\n|          |\n|         / \\ \n|\n|\n|")
}

func game_loop() {
	for hangman.tries != MAX_TRIES {
		print_elements()

	}
	print_lose_screen()
}

func main() {
	fmt.Println("Welcome to the Hangman game\n\n")

	word := generate_random_word()
	guessed := guessed_init(word)
	hangman_init(word, guessed)

	game_loop()
}
