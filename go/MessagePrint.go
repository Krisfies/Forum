package main

import "fmt"

// MessagesPrint print message.
func MessagesPrint() ([]Message, error) {

	var messages []Message

	rows, err := db.Query("SELECT * FROM `data_message` ORDER BY `data_message`.`id` ASC ")
	if err != nil {
		return nil, fmt.Errorf("MessagesPrint error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var mess Message
		if err := rows.Scan(&mess.ID, &mess.Content, &mess.Author, &mess.Date); err != nil {
			return nil, fmt.Errorf("MessagesPrint error: %v", err)
		}
		messages = append(messages, mess)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("MessagesPrint error: %v", err)
	}
	return messages, nil
}
