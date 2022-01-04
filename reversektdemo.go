package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	html :=
		`<doctype html>
        <html>
	<head>
		<title>Hello To Reversekt demo</title>
	</head>
	<body>
		<b>Welcome to reverse kt demo</b>
        <p>
            <a href="/page1">PAGE1</a> |  <a href="/page2">PAGE 2</a>
        </p>
	</body>
</html>`
	fmt.Fprintf(w, html)
}

func page1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is page 1\n")
}
func page2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is page 2\n")
}

func main() {
	// NewServeMux returns a new ServeMux.
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/page1", page1)
	mux.HandleFunc("/page2", page2)
	log.Println("Listening...")
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
