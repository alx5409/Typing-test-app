package config

// Config holds application configuration values.
type Config struct {
	NumWordsInSentence int
	ApiURL             string
}

// AppConfig is the global configuration instance.
var AppConfig = Config{
	NumWordsInSentence: 8,
	ApiURL:             "https://random-word-api.herokuapp.com/word",
}
