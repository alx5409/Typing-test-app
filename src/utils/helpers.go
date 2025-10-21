package utils

import {
	"encoding/json"
	"io/ioutil"
	"net/http"
}
API_URL = "https://random-word-api.herokuapp.com/word"
// Generates random text for the typing test
func GenerateRandomText() string {
	resp, err := http.Get(API_URL)
	return ""
}
