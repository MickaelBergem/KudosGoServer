package main

/*
   This server should handle PUT request with two parameters: the URL and the Kudo ID used.
*/

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var (
	portNumber  int
	databaseURL string
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

	if r.Method == "POST" {
		// We want to increase the corresponding kudo Button
		fmt.Printf("Increasing KudoButton with ID %s\n", kudoID)
		affectedButtons := increaseKudoButton(kudoID)
		if affectedButtons == 1 {
			fmt.Fprintf(w, "Button increased :)")
		} else {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Unknown button... :(")
		}
		return
	}

	if r.Method == "PUT" {
		// We want to create the corresponding button
		err := r.ParseForm()
		checkErr(err)

		var kudoButton KudoButton
		kudoButton.ID = kudoID
		kudoButton.URL = r.PostFormValue("URL")
		success := kudoButton.create()

		if success {
			fmt.Printf("Created new button for ID %s at URL %s\n", kudoButton.ID, kudoButton.URL)
			fmt.Fprintf(w, "Button successfully created")
		} else {
			w.WriteHeader(400)
			fmt.Printf("Error while creating the button, does the button already exist?\n")
			fmt.Fprintf(w, "Error while creating the button, does the button already exist?")
		}
		return
	}

	// Method not allowed
	w.WriteHeader(405)
}

func setUpDatabase() {
	db, err := sql.Open("sqlite3", databaseURL)

	body, err := ioutil.ReadFile("initdb.sql")
	checkErr(err)

	db.Exec(string(body))
}

func main() {

	portNumber = 8090
	databaseURL = "./kudos_count.sqlite3"

	fmt.Printf("Launching Kudos Server on port %d...\n", portNumber)

	if _, err := os.Stat(databaseURL); os.IsNotExist(err) {
		// The database does not exist, attempt to create it
		fmt.Print("Database not found, trying to initialize it...")
		setUpDatabase()
		fmt.Print(" done.\n")
	}

	// TODO: handle errors here
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+strconv.Itoa(portNumber), nil)
}
