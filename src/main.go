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

func fetchSentenceAsync() (string, error) {
	// Generate or fetch a random sentence
	fmt.Println("Fetching sentence, please wait")
	sentence, err := utils.GenerateRandomSentence()
	if err != nil {
		return "", err
	}
	return sentence, nil
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
	typingTestResult.TypingSpeed = typingTestResult.ComputeTypingSpeedWPM()
	return typingTestResult
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Typing test App")
	sentence, err := fetchSentenceAsync()
	if err != nil {
		fmt.Println("Error fetching sentence:", err)
		return
	}
	fmt.Println("Type the following sentence as fast and accurately as you can")
	fmt.Println("\n>>>", sentence)

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
	typingTestResult := buildTest(sentence, userInput, start, end)

	// Print the results using the typing package
	typing.ShowResults(typingTestResult)
}
