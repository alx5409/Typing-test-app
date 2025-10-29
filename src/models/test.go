package models

import (
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
	if len(t.TextTyped) == 0 {
		return 0
	}
	return 1 - float32(t.NumberErrors)/float32(len(t.TextTyped))
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
	return (float32(len(t.TextTyped)) / duration) * 60
}

// Computes the user's words per minute (WPM) counting words
// as sequences separated by spaces.
// Returns 0 i fthe duration is zero.
func (t *TypeTest) ComputeTypingSpeedWPM() float32 {
	duration := t.ComputeTestTime()
	if duration == 0 {
		return 0
	}
	wordCount := len(strings.Fields(t.TextTyped))
	if wordCount == 0 {
		return 0
	}
	return (float32(wordCount) / duration) * 60
}
