package main

import "fmt"

func AddTopic(top Data_Topic) (int64, error) {
	result, err := db.Exec("INSERT INTO data_topic (Name) VALUES (?)", top.Name)
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
