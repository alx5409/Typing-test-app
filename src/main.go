package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	typing "Typing-test-app/src/handlers"
	models "Typing-test-app/src/models"
	utils "Typing-test-app/src/utils"
)

func FetchSentenceAsync() string {
	// Generate or fetch a random sentence
	sentenceCh := make(chan string)
	go func() {
		sentenceCh <- utils.GenerateRandomSentence()
	}()
	fmt.Println("Fetching sentence, please wait...")
	return <-sentenceCh
}

func main() {
	fmt.Println("Welcome to the Typing test App")
	sentence := FetchSentenceAsync()
	fmt.Println("Type the following sentence as fast and accurately as you can")
	fmt.Println("\n>>>", sentence)
	// fmt.Println("\nPress enter when you are ready to start")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	start := time.Now()
	fmt.Print("Your input: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	elapsed := time.Since(start)

	// Prepare the test struct
	test := &models.TypeTest{
		TextToType:   sentence,
		TextTyped:    userInput,
		StartTime:    start,
		EndTime:      start.Add(elapsed),
		NumberErrors: utils.CountErrors(sentence, userInput),
	}
	test.TypingSpeed = test.ComputeTypingSpeedWPM()

	// Print the results using the typing package
	typing.ShowResults(test)
}
