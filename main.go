package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"io/ioutil"
)

type buzzwords struct {
	Adjectives []string `json:"adjectives"`
	Nouns      []string `json:"nouns"`
}

var (
	words buzzwords
)

func main() {
	words = readJSONFile("buzzwords.json")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/words", wordsHandler)

	log.Fatal(http.ListenAndServe(":3001", nil))

}

func readJSONFile(filename string) buzzwords {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(raw, &words)
	return words
}

func responseWithJSON(w http.ResponseWriter, body interface{}, code int) {
	result, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(result)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	adjective := words.Adjectives[rand.Intn(len(words.Adjectives))]
	noun := words.Nouns[rand.Intn(len(words.Nouns))]
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(adjective + " " + noun))
}

func wordsHandler(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, words, 200)
}
