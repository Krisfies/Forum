package main

import "fmt"

func AddMessage(mess Data_Message) (int64, error) {
	// Crée un message dans la base de donnée
	result, err := db.Exec("INSERT INTO data_message (Content, Author, Date, Topic_ID) VALUES (?, ?, ?, ?)", mess.Content, mess.Author, mess.Date, mess.Topic_ID)
	// Gère les erreur pour envoyer le message
	if err != nil {
		return 0, fmt.Errorf("AddMessage: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddMessage: %v", err)
	}
	return id, nil
}
