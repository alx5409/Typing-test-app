package config

import (
	"os"

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
	// AppConfig is the global configuration instance.
	var AppConfig = Config{
		NumWordsInSentence: 30,
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
