package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// A KudoButton is a Kudo button on a page
type KudoButton struct {
	URL       string
	ID        string
	KudoCount int
}

//
func (kudoButton *KudoButton) create() {

	db, err := sql.Open("sqlite3", databaseUrl)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO kudos (KudoID, KudoCount, URL) VALUES (?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(kudoButton.ID, kudoButton.KudoCount, kudoButton.URL)
	checkErr(err)

	db.Close()
}

// getCurrentCount returns the current number of Kudos for the given ID
func getCurrentCount(kudoID string) int {
	return getKudoButton(kudoID).KudoCount
}

func getKudoButton(kudoID string) *KudoButton {

	db, err := sql.Open("sqlite3", databaseUrl)
	checkErr(err)
	stmt, err := db.Prepare("SELECT KudoCount, URL FROM kudos WHERE KudoID = ?")
	checkErr(err)
	rows, err := stmt.Query(kudoID)
	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		var kudoButton KudoButton

		err = rows.Scan(
			&kudoButton.KudoCount,
			&kudoButton.URL,
		)

		checkErr(err)

		db.Close()

		fmt.Printf("Found the KudoButton, count=%d\n", kudoButton.KudoCount)
		return &kudoButton
	}

	db.Close()

	fmt.Printf("Unable to find a kudo with ID=%s, return empty button...\n", kudoID)
	return &KudoButton{KudoCount: 0}
}

// Increase the KudoButton with kudoID and return 1 if there was a matching KudoButton
func increaseKudoButton(kudoID string) int64 {
	db, err := sql.Open("sqlite3", databaseUrl)
	checkErr(err)
	stmt, err := db.Prepare("UPDATE kudos SET KudoCount = KudoCount + 1 WHERE KudoID = ?")
	checkErr(err)
	res, err := stmt.Exec(kudoID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	db.Close()

	return affect
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
