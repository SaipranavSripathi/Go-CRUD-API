package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, from Sai Pranav - 1st Go API")
}
func dataPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Make sure you reached getData route")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getData", dataPage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
