package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration values.
type Config struct {
	NumWordsInSentence int
	ApiURL             string
	SupportedLanguages map[string]string // key: language name, value: API code
}

func Initialize() Config {
	godotenv.Load()
	numStr := os.Getenv("NUMBER_OF_WORDS_PER_SENTENCE")
	numWords, err := strconv.Atoi(numStr)
	defaultNumWords := 30
	if err != nil {
		numWords = defaultNumWords // default value if conversion fails
	}
	// AppConfig is the global configuration instance.
	var AppConfig = Config{
		NumWordsInSentence: numWords,
		// ApiURL:             "https://random-word-api.vercel.app/api",
		ApiURL: os.Getenv("API_URL"),
		SupportedLanguages: map[string]string{
			"English":              "",
			"Spanish":              "es",
			"Italian":              "it",
			"German":               "de",
			"French":               "fr",
			"Chinese":              "zh",
			"Brazilian Portuguese": "pt-br",
		},
	}
	return AppConfig
}

func (c *Config) Print() {
	println("Current Configuration:")
	println("NumWordsInSentence:", c.NumWordsInSentence)
	println("ApiURL:", c.ApiURL)
}
