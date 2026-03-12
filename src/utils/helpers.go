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
	"log"
	"net/http"
	"strings"
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
func GenerateRandomSentence() (string, error) {
	numWords := config.AppConfig.NumWordsInSentence
	completeURL := config.AppConfig.ApiURL + "?number=" + fmt.Sprintf("%d", numWords)
	resp, err := http.Get(completeURL)
	if err != nil {
		log.Default().Printf("Error fetching sentence: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Default().Printf("Error reading response body: %v", err)
		return "", err
	}
	var rawWords []string
	err = json.Unmarshal(body, &rawWords)
	if err != nil {
		log.Default().Printf("Error parsing JSON response: %v", err)
		return "", err
	}
	validWords := []string{}
	for _, w := range rawWords {
		word := strings.TrimSpace(w)
		word = strings.Trim(word, "\"") // remove quotes
		if word == "" || strings.Contains(word, " ") {
			continue
		}
		validWords = append(validWords, word)
	}
	return strings.Join(validWords, " "), nil
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

// iterative, rune-aware Levenshtein (DP) or edit distance implementation
func Levenshtein(a, b string) int {
	ar := []rune(a)
	br := []rune(b)
	if len(ar) == 0 {
		return len(br)
	}
	if len(br) == 0 {
		return len(ar)
	}

	prev := make([]int, len(br)+1)
	cur := make([]int, len(br)+1)
	for j := 0; j <= len(br); j++ {
		prev[j] = j
	}
	for i := 1; i <= len(ar); i++ {
		cur[0] = i
		for j := 1; j <= len(br); j++ {
			cost := 0
			if ar[i-1] != br[j-1] {
				cost = 1
			}
			ins := cur[j-1] + 1
			del := prev[j] + 1
			sub := prev[j-1] + cost
			m := ins
			if del < m {
				m = del
			}
			if sub < m {
				m = sub
			}
			cur[j] = m
		}
		copy(prev, cur)
	}
	return prev[len(br)]
}

// Calculates the similarity between two words based on their edit distance.
func WordsMatchAmount(target, typed string) float32 {
	t := strings.TrimSpace(target)
	p := strings.TrimSpace(typed)
	if t == p {
		return 1.0
	}
	dist := Levenshtein(t, p)
	return 1 - float32(dist)/float32(max(len([]rune(t)), len([]rune(p))))
}
