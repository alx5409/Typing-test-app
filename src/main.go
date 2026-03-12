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

type fetchResult struct {
	sentence string
	err      error
}

func fetchSentenceAsync() chan fetchResult {
	// Create a channel to receive the result of the fetch operation
	ch := make(chan fetchResult, 1)
	// Generate or fetch a random sentence
	fmt.Println("Fetching sentence, please wait")
	go func() {
		sentence, err := utils.GenerateRandomSentence()
		ch <- fetchResult{sentence: sentence, err: err}
	}()
	return ch
}

func waitForStart(reader *bufio.Reader) error {
	fmt.Println("\nPress enter when you are ready to start")
	_, err := reader.ReadBytes('\n')
	return err
}

func readTypedInput(reader *bufio.Reader) (string, time.Time, time.Time, error) {
	fmt.Print("Your input: ")
	start := time.Now()
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", time.Time{}, time.Time{}, err
	}
	end := time.Now()
	return strings.TrimSpace(userInput), start, end, nil
}

func buildTest(sentence, userInput string, start, end time.Time) *models.TypeTest {
	typingTestResult := &models.TypeTest{
		TextToType:   sentence,
		TextTyped:    userInput,
		StartTime:    start,
		EndTime:      end,
		NumberErrors: utils.CountErrors(sentence, userInput),
	}
	return typingTestResult
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Typing test App")
	resCh := fetchSentenceAsync()
	result := <-resCh
	if result.err != nil {
		fmt.Println("Error fetching sentence:", result.err)
		return
	}
	fmt.Println("Type the following sentence as fast and accurately as you can")
	fmt.Println("\n>>>", result.sentence)

	if err := waitForStart(reader); err != nil {
		fmt.Println("Error waiting for start:", err)
		return
	}

	userInput, start, end, err := readTypedInput(reader)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// Prepare the typingTestResult struct
	typingTestResult := buildTest(result.sentence, userInput, start, end)

	// Print the results using the typing package
	typing.ShowResults(typingTestResult)
}
