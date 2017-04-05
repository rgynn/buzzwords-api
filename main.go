package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/hajhatten/buzzwords"
)

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/buzzword", verbsHandler)
	http.HandleFunc("/verb", verbsHandler)
	http.HandleFunc("/suffix", suffixHandler)
	http.HandleFunc("/verbsuffix", verbsAndSuffixHandler)

	var port string
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":3001"
	}

	log.Fatal(http.ListenAndServe(port, nil))

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

func responseWithText(w http.ResponseWriter, body string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(body))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func buzzwordsHandler(w http.ResponseWriter, r *http.Request) {
	responseWithText(w, buzzwords.BuzzWords(), 200)
}

func verbsHandler(w http.ResponseWriter, r *http.Request) {
	responseWithText(w, buzzwords.WithVerb(), 200)
}

func suffixHandler(w http.ResponseWriter, r *http.Request) {
	responseWithText(w, buzzwords.WithSuffix(), 200)
}

func verbsAndSuffixHandler(w http.ResponseWriter, r *http.Request) {
	responseWithText(w, buzzwords.WithVerbAndSuffix(), 200)
}
