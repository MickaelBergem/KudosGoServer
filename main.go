package main

/*
   This server should handle PUT request with two parameters: the URL and the Kudo ID used.
*/

import (
	"fmt"
	"net/http"
	"strconv"
)

var (
	portNumber  int
	databaseUrl string
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got a %s request on %s\n", r.Method, r.URL)

	// We expect path to be like /kudo-id-string
	kudoID := r.URL.Path[1:]

	if r.Method == "GET" {
		// Fetch the corresponding button
		fmt.Printf("Fetching KudoButton with ID %s\n", kudoID)
		kudoButton := getKudoButton(kudoID)

		fmt.Fprintf(w, "Hey, this button has %d kudos!", kudoButton.KudoCount)
		return
	}

	if r.Method == "PUT" {
		// We want to increase the corresponding kudo Button
		fmt.Printf("Increasing KudoButton with ID %s\n", kudoID)
		affectedButtons := increaseKudoButton(kudoID)
		if affectedButtons == 1 {
			fmt.Fprintf(w, "Button increased :)")
		} else {
			fmt.Fprintf(w, "Unknown button... :(")
		}
		return
	}

	// Method not allowed
	w.WriteHeader(405)
}

func main() {

	portNumber = 8090
	databaseUrl = "./kudos_count.sqlite3"

	fmt.Printf("Launching Kudos Server...\n")

	// TODO: handle errors here
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+strconv.Itoa(portNumber), nil)
}
