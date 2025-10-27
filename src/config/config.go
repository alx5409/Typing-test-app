package config

// Config holds application configuration values.
type Config struct {
	NumWordsInSentence int
	ApiURL             string
	SupportedLanguages map[string]string // key: language name, value: API code
}

// AppConfig is the global configuration instance.
var AppConfig = Config{
	NumWordsInSentence: 8,
	ApiURL:             "https://random-word-api.herokuapp.com/word",
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
