package main

import "fmt"

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func AddMessage(mess Message) (int64, error) {
	result, err := db.Exec("INSERT INTO message (Content, Author, Date) VALUES (?, ?, ?)", mess.Content, mess.Author, mess.Date)
	fmt.Println(result)
	if err != nil {
		return 0, fmt.Errorf("AddMessage: %v", err)
	}
	id, err := result.LastInsertId()
	fmt.Println(id)
	if err != nil {
		return 0, fmt.Errorf("AddMessage: %v", err)
	}
	return id, nil
}
