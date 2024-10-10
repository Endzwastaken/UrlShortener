package endpoint

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Service interface {
	GenerateShortKey() string
	Insert(string, string) error
	Get(string) (string, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Form(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>URL Shortener</title>
		</head>
		<body>
			<h2>URL Shortener</h2>
			<form method="post" action="/s">
				<input type="url" name="url" placeholder="Enter a URL" required>
				<input type="submit" value="Shorten">
			</form>
		</body>
		</html>
	`)
}

func (e *Endpoint) Shorting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	shortKey := e.s.GenerateShortKey()
	e.s.Insert(shortKey, originalURL)

	shortenedURL := fmt.Sprintf("http://localhost:8080/%s", shortKey)

	w.Header().Set("Content-Type", "text/html")
	responseHTML := fmt.Sprintf(`
        <h2>URL Shortener</h2>
        <p>Original URL: %s</p>
        <p>Shortened URL: <a href="%s">%s</a></p>
		<a href="http://localhost:8080/">Main page</a></p>
    `, originalURL, shortenedURL, shortenedURL)
	fmt.Fprint(w, responseHTML)
	fmt.Println("-----------------------------------")
	fmt.Printf("%s %s%s\nRequest: %s\nResponse: %s\n", r.Method, r.Host, r.Pattern, originalURL, shortenedURL)
	fmt.Println("-----------------------------------")
}

func (e *Endpoint) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortKey := vars["shortKey"]
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	originalURL, found := e.s.Get(shortKey)
	if found != nil {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
	fmt.Println("-----------------------------------")
	fmt.Printf("%s Request: %s%s\nResponse: %s\n", r.Method, r.Host, r.URL, originalURL)
	fmt.Println("-----------------------------------")
}
