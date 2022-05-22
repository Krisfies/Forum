package main

import "fmt"

func AddTopic(top Data_Topic) (int64, error) {
	Aide, Bug, Boss, Lore := IsTagInTopic(top)
	result, err := db.Exec("INSERT INTO data_topic (Name, IsAide, IsBug, IsBoss, IsLore) VALUES (?,?,?,?,?)", top.Name, Aide, Bug, Boss, Lore)
	fmt.Println(result)
	if err != nil {
		return 0, fmt.Errorf("AddTopic: %v", err)
	}
	id, err := result.LastInsertId()
	fmt.Println(id)
	if err != nil {
		return 0, fmt.Errorf("AddTopic: %v", err)
	}
	return id, nil
}
