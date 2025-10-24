package typing

import (
	"Typing-test-app/src/models"
	"Typing-test-app/src/utils"
	"fmt"
	"time"
)

// StartTest initializes a new typing test, generates the text, and sets the start time.
func StartTest() *models.TypeTest {
	test := &models.TypeTest{
		TextToType: utils.GenerateRandomText(),
		StartTime:  time.Now(),
	}
	fmt.Println("Typing test started. Type the following text:")
	fmt.Println(test.TextToType)
	return test
}

// EndTest finalizes the typing test, sets the end time, and processes the results.
func EndTest(test *models.TypeTest, typedText string) {
	test.TextTyped = typedText
	test.EndTime = time.Now()
	test.NumberErrors = utils.CountErrors(test.TextToType, test.TextTyped)
	test.TypingSpeed = test.ComputeTypingSpeedWPM()
}

// ShowResults displays the results of the typing test.
func ShowResults(test *models.TypeTest) {
    fmt.Printf("Accuracy: %.2f%%\n", test.ComputeAccuracy()*100)
    fmt.Printf("Keystrokes per minute (PPM): %.2f\n", test.ComputeTypingSpeedPPM())
    fmt.Printf("Words per minute (WPM): %.2f\n", test.ComputeTypingSpeedWPM())
    fmt.Printf("Errors: %d\n", test.NumberErrors)
    fmt.Printf("Test duration: %.2f seconds\n", test.ComputeTestTime())
}