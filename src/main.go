package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"Typing-test-app/src/utils"
)

func main() {
	fmt.Println("Welcome to the Typing test App")
	fmt.Println("Type the following sentence as fast and accurately as you can")

	// Generate or fetch a random sentence
	sentence := utils.GenerateRandomSentence()
	fmt.Println("\n>>>", sentence)
	fmt.Println("\nPress enter when you are ready to start")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	start := time.Now()
	fmt.Print("Your input: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	elapsed := time.Since(start)

	// Calculate results
	errors := utils.CountErrors(sentence, userInput)
	wpm := utils.CalculateWPM(sentence, elapsed)

	// Print the results
	fmt.Printf("\nResults:\n")
	fmt.Printf("Time: %.2f seconds\n", elapsed.Seconds())
	fmt.Printf("Errors: %d\n", errors)
	fmt.Printf("Speed: %.2f WPM\n", wpm)
}
