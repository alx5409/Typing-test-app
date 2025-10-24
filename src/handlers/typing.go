package typing

import (
	"Typing-test-app/src/models"
	"Typing-test-app/src/utils"
	"fmt"
	"time"
)

func StartTest() *models.TypeTest {
	test := &models.TypeTest{
		TextToType: utils.GenerateRandomText(),
		StartTime:  time.Now(),
	}
	fmt.Println("Typing test started. Type the following text:")
	fmt.Println(test.TextToType)
	return test
}
