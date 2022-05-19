package main

import "fmt"

// MessagesPrint print message.
func MessagesPrint(ID int) ([]Data_Message, error) {

	var messages []Data_Message

	rows, err := db.Query("SELECT * FROM `data_message` WHERE `Topic_ID` = ? ORDER BY `data_message`.`ID` ASC", ID)
	if err != nil {
		return nil, fmt.Errorf("MessagesPrint error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var mess Data_Message
		if err := rows.Scan(&mess.ID, &mess.Content, &mess.Author, &mess.Date, &mess.Topic_ID); err != nil {
			return nil, fmt.Errorf("MessagesPrint error: %v", err)
		}
		messages = append(messages, mess)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("MessagesPrint error: %v", err)
	}
	return messages, nil
}
