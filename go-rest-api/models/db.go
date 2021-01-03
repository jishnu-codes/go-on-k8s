package models

import (
	"database/sql"
)

// The global variable below handles the database connetion
var DB *sql.DB

type Message struct {
	Greeting string
	Year     int
}

// The Hello() returns a slice of messages from the messages table
func Hello() (msgs []Message, err error) {
	// Calling the Query() fucntion on the global variable.
	rows, err := DB.Query("SELECT * FROM mygreeting")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//var msgs []Message
	var msg Message
	for rows.Next() {

		err := rows.Scan(&msg.Greeting, &msg.Year)
		if err != nil {
			return nil, err
		}

		msgs = append(msgs, msg)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return msgs, nil
}
