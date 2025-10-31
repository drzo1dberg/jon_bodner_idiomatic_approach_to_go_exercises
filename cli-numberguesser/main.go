package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var highScores = map[string]int{}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the number guessing game!")
	fmt.Println("Im thinking of a number between 1 and 200.")
	fmt.Println("Type a number to guess. -> Type 'hint' for ONE clue per round. -> type 'quit' to exit")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		difficulty, attempts := chooseDifficulty(reader)
		used, took, won := playRound(reader, difficulty, attempts)
	}
}

func chooseDifficulty(reader *bufio.Reader) (string, int) {
	fmt.Println("Please select a difficulty lvl:")
	fmt.Println("1. Easy		(10 chances)")
	fmt.Println("2. Medium		(5 chances)")
	fmt.Println("3. Hard		(3 chances)")
	for {
		fmt.Print("Enter your choice (1/2/3 or easy/medium/hard): ")
		line, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(strings.ToLower(line))

		switch choice {
		case "1", "easy":
			fmt.Println("You selected easy.")
			return "easy", 10
		case "2", "medium":
			fmt.Println("You selected medium.")
			return "medium", 5
		case "3", "hard":
			fmt.Println("You selected hard.")
			return "hard", 3
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func playRound(reader *bufio.Reader, difficulty string, attempts int) (used int, took time.Duration, won bool) {
	target := rand.Intn(200) + 1
	start := time.Now()

	hintsLeft := 1

	fmt.Printf("\nLet's start! You have %d chances. Good luck!\n", attempts)

	for attempt := 1; attempt <= attempts; attempt++ {
		fmt.Printf("Attempt %d/%d - Enter your guess: ", attempt, attempts)
		line, _ := reader.ReadString('\n')
		input := strings.TrimSpace(strings.ToLower(line))

		if input == "quit" {
			fmt.Println("Quitting the game")
			return attempt - 1, time.Since(start), false
		}

		if input == "hint" {
			if hintsLeft <= 0 {
				fmt.Println("No hints left this round.")
				attempt--
				continue
			}
			hintsLeft--
			printHint(target)
			attempt--
			continue
		}

		guess, err := strconv.Atoi(input)
		if err != nil || guess < 1 || guess > 200 {
			fmt.Println("Please enter a whole number between 1 and 200, or 'hint'.")
			attempt--
			continue
		}

		if guess == target {
			return attempt, time.Since(start), true
		}
		if guess < target {
			fmt.Printf("WRONG! The Number is greater than %d.\n", guess)
		} else {
			fmt.Printf("WRONG! The number is less than %d.\n", guess)
		}
	}
	return attempts, time.Since(start), false
}

func printHint(target int) {
	parity := "even"
	if target%2 != 0 {
		parity = "odd"
	}

	bucket := "1-25"
	switch {
	case target <= 25:
		bucket = "1-25"
	case target <= 50:
		bucket = "26-50"
	case target <= 75:
		bucket = "51-75"
	case target <= 100:
		bucket = "76-100"
	case target <= 125:
		bucket = "101-125"
	case target <= 150:
		bucket = "126-150"
	case target <= 175:
		bucket = "151-175"
	default:
		bucket = "176-200"
	}
	fmt.Printf(" Hint: Its an %s number in the %s range. \n", parity, bucket)
}

func updateHighScore()
