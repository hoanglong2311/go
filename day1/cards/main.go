package main

import (
	"fmt"
	"net/http"
)

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Diamonds"
	fmt.Println(card)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)

}
