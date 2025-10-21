package models

import "time"

type TypeTest struct {
	Id           int64
	TextToType   string
	TextTyped    string
	StartTime    time.Time
	EndTime      time.Time
	NumberErrors int
	TypingSpeed  float32
}

func (t *TypeTest) ComputeAccuracy() float32 {
	if len(t.TextTyped) == 0 {
		return 0
	}
	return 1 - float32(t.NumberErrors)/float32(len(t.TextTyped))
}

func (t *TypeTest) ComputeTestTime() float32 {
	duration := t.EndTime.Sub(t.StartTime).Seconds()
	return float32(duration)
}

func (t *TypeTest) ComputeTypingSpeedPPM() float32 {
	duration := t.ComputeTestTime()
	if duration == 0 {
		return 0
	}
	return (float32(len(t.TextTyped)) / duration) * 60
}

func (t *TypeTest) ComputeTypingSpeedWPM() float32 {
	duration := t.ComputeTestTime()
	if duration == 0 {
		return 0
	}
	wordCount := 1
	for _, c := range t.TextTyped {
		if c == ' ' {
			wordCount++
		}
	}
	return (float32(wordCount) / duration) * 60
}
