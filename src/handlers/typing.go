package typing

import (
	"Typing-test-app/src/models"
	"Typing-test-app/src/utils"
	"fmt"
	"time"
)

// StartTest initializes a new typing test, generates the text to type, and sets the start time.
// It prints the text that the user should type and returns a pointer to the initialized TypeTest struct.
func StartTest() *models.TypeTest {
	test := &models.TypeTest{
		TextToType: utils.GenerateRandomText(),
		StartTime:  time.Now(),
	}
	fmt.Println("Typing test started. Type the following text:")
	fmt.Println(test.TextToType)
	return test
}

// EndTest finalizes the typing test by setting the end time, saving the user's input,
// counting errors, and calculating the typing speed (WPM).
func EndTest(test *models.TypeTest, typedText string) {
	test.TextTyped = typedText
	test.EndTime = time.Now()
	test.NumberErrors = utils.CountErrors(test.TextToType, test.TextTyped)
	test.TypingSpeed = test.ComputeTypingSpeedWPM()
}

// ShowResults displays the results of the typing test, including accuracy,
// keystrokes per minute (PPM), words per minute (WPM), number of errors, and test duration.
func ShowResults(test *models.TypeTest) {
	fmt.Printf("Accuracy: %.2f%%\n", test.ComputeAccuracy()*100)
	fmt.Printf("Keystrokes per minute (PPM): %.2f\n", test.ComputeTypingSpeedPPM())
	fmt.Printf("Words per minute (WPM): %.2f\n", test.ComputeTypingSpeedWPM())
	fmt.Printf("Errors: %d\n", test.NumberErrors)
	fmt.Printf("Test duration: %.2f seconds\n", test.ComputeTestTime())
}