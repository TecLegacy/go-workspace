package main

import "fmt"

func main() {
	testBadWord()
}

//* Index of first badword
func testBadWord() {
	badWord := []string{
		"stupid",
		"idiot",
		"duffer",
	}
	message := []string{"I", "Love", "you", "stupid"}

	index := indexOfFirstBadWord(message, badWord)

	fmt.Printf("The index of badword %d", index+1)
}

func indexOfFirstBadWord(msg, badWords []string) int {
	// On^2 time complexity
	for i, word := range msg {
		for _, badWord := range badWords {
			if badWord == word {
				return i
			}
		}
	}
	return -1
}

//* fizzbuzz classic problem
func fizzBud(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
