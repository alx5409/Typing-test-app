// Package utils provides utility functions for the Typing-test-app.
// It includes helpers for generating random text, counting typing errors,
// normalizing text, generating sentences, handling language options,
// caching random text, and calculating words per minute (WPM).
//
// Functions in this package are used throughout the application to support
// the core typing test logic and ensure consistent text processing.

package utils

import (
	config "Typing-test-app/src/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Returns the maximum value
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Generates random text for the typing test
func GenerateRandomText() string {
	resp, err := http.Get(config.AppConfig.ApiURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var words []string
	if err := json.Unmarshal(body, &words); err != nil {
		return ""
	}

	// Joins the words into a single string separated by spaces
	// skipping any word that cointais spaces
	text := ""
	for _, word := range words {
		if strings.Contains(word, " ") {
			continue // skip words with spaces
		}
		if text != "" {
			text += " " // Joins the words with spaces
		}
		text += word
	}
	return text
}

// Counts the number of errors between expected and typed text
func CountErrors(expected, typed string) int {
	expectedRunes := []rune(expected)
	typedRunes := []rune(typed)
	errors := 0
	// Determine the maximum length to compare both strings fully
	maxLen := max(len(expectedRunes), len(typedRunes))

	// Compare each rune in both strings
	for i := 0; i < maxLen; i++ {
		var e, t rune
		// Get the rune from expected if it is within bounds
		if i < len(expectedRunes) {
			e = expectedRunes[i]
		}
		// Get the rune from typed if it is within bounds
		if i < len(typedRunes) {
			t = typedRunes[i]
		}
		// If one string is shorter or runes differ, count as error
		if i >= len(expectedRunes) || i >= len(typedRunes) || e != t {
			errors++
		}
	}
	return errors
}

// Normalizes the input text (e.g., removes special characters, converts to lowercase)
func NormalizeText(text string) string {
	text = strings.ToLower(text)
	var builder strings.Builder
	for _, r := range text {
		// Removes any character that is not lowercase ASCII letter(a - z) or number (0 - 9)
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// Generates a random sentence or paragraph for the typing test
func GenerateRandomSentence() string {
	numWords := config.AppConfig.NumWordsInSentence
	// words := []string{}
	resp, err := http.Get(config.AppConfig.ApiURL + "?words=" + fmt.Sprintf("%d", numWords))

	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var words []string
	if err := json.Unmarshal(body, &words); err != nil {
		return ""
	}
	// Filter out words with spaces (just in case)
	validWords := []string{}
	for _, word := range words {
		word = strings.TrimSpace(word)
		if word == "" || strings.Contains(word, " ") {
			continue
		}
		validWords = append(validWords, word)
	}
	return strings.Join(validWords, " ")
}

// Gets random words in a specified language (if API supports)
// Appends the language code as a query parameter (?lang=xx) to the API URL.
func GetRandomTextWithLanguage(language string) string {
	apiURL := config.AppConfig.ApiURL
	if language != "" {
		sep := "?"
		if strings.Contains(apiURL, "?") {
			sep = "&"
		}
		apiURL = apiURL + sep + "lang=" + language
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var words []string
	if err := json.Unmarshal(body, &words); err != nil {
		return ""
	}

	text := ""
	for _, word := range words {
		if strings.Contains(word, " ") {
			continue // skip words with spaces
		}
		if text != "" {
			text += " "
		}
		text += word
	}
	return text
}

// Returns cached random text or fetches new text if cache is empty.
// This simple implementation uses a package-level variable as cache.
var cachedText string

func GetCachedRandomText() string {
	if cachedText != "" {
		return cachedText
	}
	cachedText = GenerateRandomText()
	return cachedText
}

// Computes words per minuto given the sentence and elapsed time.
// Words are cunted as sequences separated by spaces
func CalculateWPM(sentence string, elapsed time.Duration) float64 {
	sentence = strings.TrimSpace(sentence)
	if sentence == "" {
		return 0 // No words to count if the sentence is empty
	}
	words := strings.Fields(sentence) // separates the sentence by spaces in different strings
	wordCount := len(words)           // counts how many strings there are
	minutes := elapsed.Minutes()
	if minutes == 0 {
		return 0
	}
	return float64(wordCount) / minutes
}
