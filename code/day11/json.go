package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func decode(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Fprintf(w, "%s %s is %d years old!\n", user.FirstName, user.LastName, user.Age)
}

func encode(w http.ResponseWriter, r *http.Request) {
	john := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	json.NewEncoder(w).Encode(john)
}

func main() {
	http.HandleFunc("/decode", decode)
	http.HandleFunc("/encode", encode)

	http.ListenAndServe(":8080", nil)
}
