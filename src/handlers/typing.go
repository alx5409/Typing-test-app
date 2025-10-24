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
