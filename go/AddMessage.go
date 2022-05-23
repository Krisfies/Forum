package main

import "fmt"

// AddMessage adds the specified message to the database,
// returning the message ID of the new entry
func AddMessage(mess Data_Message) (int64, error) {
	result, err := db.Exec("INSERT INTO data_message (Content, Author, Date, Topic_ID) VALUES (?, ?, ?, ?)", mess.Content, mess.Author, mess.Date, mess.Topic_ID)
	if err != nil {
		return 0, fmt.Errorf("AddMessage: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddMessage: %v", err)
	}
	return id, nil
}
