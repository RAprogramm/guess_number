// Simple game. Player must guess a number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // get date and time as integer
	rand.Seed(seconds)           // init random numbers generator
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin) // create for read data from keyboard

	success := false // by default you will see message about loose
	for guesses := 0; guesses < 5; guesses++ {
		fmt.Println("You have", 5-guesses, "guesses left.")

		fmt.Println("Make a guess: ")         // get a number from player
		input, err := reader.ReadString('\n') // read data before "Enter" pressed
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  // remove new line symbol
		guess, err := strconv.Atoi(input) // from "string" to "int"
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("No. My number is MORE than", guess)
		} else if guess > target {
			fmt.Println("Oops. My number is LESS than", guess)
		} else {
			success = true // remove message about loose
			fmt.Println("Good job! You guess it!")
			break
		}
	}
	if !success { // inform about "target" if "success" equal "false"
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}
