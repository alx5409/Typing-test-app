package utils

// TODO:
// - Implement a function to count erros at typing
// - Implement a function to normalize text (remove extra spaces, lowercase, etc.).
// - Implement support for generating sentences or paragraphs instead of single words.
// - Implement caching or offline fallback if the API is unavailable.
// - Implement language selection for random words (if API supports).

import (
	"encoding/json"
	"io"
	"net/http"
)

const API_URL = "https://random-word-api.herokuapp.com/word"

// Generates random text for the typing test
func GenerateRandomText() string {
	resp, err := http.Get(API_URL)
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
	text := ""
	for i, word := range words {
		if i > 0 {
			text += " "
		}
		text += word
	}
	return text
}

// Counts the number of errors between expected and typed text
func CountErrors(expected, typed string) int {
	// TODO: implement error counting logic
	return 0
}

// Normalizes the input text (e.g., removes extra spaces, converts to lowercase)
func NormalizeText(text string) string {
	// TODO: implement normalization logic
	return text
}

// Generates a random sentence or paragraph for the typing test
func GenerateRandomSentence() string {
	// TODO: implement sentence generation logic
	return ""
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
