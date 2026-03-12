package models

import (
	helper "Typing-test-app/src/utils"
	"strings"
	"time"
)

// TypeTest represents a typing test performed by the user
type TypeTest struct {
	Id           int64     // Unique identifier for the test
	TextToType   string    // Text the user must type
	TextTyped    string    // Text entered by the user
	StartTime    time.Time // Time when the test starts
	EndTime      time.Time // Time when the test ends
	NumberErrors int       // Numbers of errors made
	TypingSpeed  float32   // Typing speed of the test
}

func (t *TypeTest) ComputeAccuracy() float32 {
	if len(t.TextToType) == 0 {
		return 0
	}
	result := 1 - float32(t.NumberErrors)/float32(len(t.TextToType))
	if result < 0 || result > 1 {
		return 0
	}
	return result
}

// Computes the total test time in seconds as float32.
func (t *TypeTest) ComputeTestTime() float32 {
	duration := t.EndTime.Sub(t.StartTime).Seconds()
	return float32(duration)
}

// Computes the user's keystrokes per minutes (PPM).
// Returns 0 if the duration is 0.
func (t *TypeTest) ComputeTypingSpeedPPM() float32 {
	duration := t.ComputeTestTime()
	if duration == 0 {
		return 0
	}
	charsTyped := len([]rune(t.TextTyped))
	return (float32(charsTyped) / duration) * 60
}

// Computes the user's words per minute (WPM) counting words
// as sequences separated by spaces.
// Returns 0 if the duration is zero.
func (t *TypeTest) ComputeTypingSpeedWPM() float32 {
	duration := t.ComputeTestTime()
	if duration <= 0 {
		return 0
	}

	targetWords := strings.Fields(t.TextToType)
	typedWords := strings.Fields(t.TextTyped)

	sumMatch := float32(0)
	for i := 0; i < min(len(targetWords), len(typedWords)); i++ {
		sumMatch += helper.WordsMatchAmount(targetWords[i], typedWords[i])
	}
	minutes := duration / 60
	if minutes == 0 {
		return 0
	}
	return (sumMatch / minutes)
}
