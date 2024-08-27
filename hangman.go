import (
	"fmt"
)

const int MAX_TRIES = 6

type Hangman struct {
	word    string
	guessed string
	tries   int
}

Hangman hangman

func main() {
	fmt.Println("Welcome to the Hangman game\n\n")

}
