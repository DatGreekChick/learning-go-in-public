package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func logRequest(w http.ResponseWriter, request string) {
	_, err := fmt.Fprintln(w, request)
	if err != nil {
		return
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	logRequest(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	logRequest(w, "bar")
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
