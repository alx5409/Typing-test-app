package utils

// TODO:
// - Implement a function to count erros at typing
// - Implement a function to normalize text (remove extra spaces, lowercase, etc.).
// - Implement support for generating sentences or paragraphs instead of single words.
// - Implement caching or offline fallback if the API is unavailable.
// - Implement language selection for random words (if API supports).

import (
	config "Typing-test-app/src/config"
	"encoding/json"
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
	words := []string{}
	for len(words) < numWords {
		word := GenerateRandomText()
		word = strings.TrimSpace(word)
		// Skip empty words
		if word == "" {
			continue
		}
		words = append(words, word)
	}
	return strings.Join(words, " ")
}

// Gets random words in a specified language (if API supports)
func GetRandomTextWithLanguage(language string) string {
	// TODO: implement language support logic
	return ""
}

// Returns cached random text or fetches new text if cache is empty
func GetCachedRandomText() string {
	// TODO: implement caching logic
	return ""
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
