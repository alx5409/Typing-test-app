package typing

import (
	"Typing-test-app/src/models"
	"Typing-test-app/src/utils"
	"fmt"
	"time"
)

// Dummy function to keep imports
func InitTypingTest() {
	fmt.Println("Typing test initialized at", time.Now())
	_ = models.TypeTest{}
	_ = utils.GenerateRandomText()
}
