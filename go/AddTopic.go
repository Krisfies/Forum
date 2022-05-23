package main

import "fmt"

func AddTopic(top Data_Topic) (int64, error) {
	Aide, Bug, Boss, Lore := IsTagInTopic(top)
	result, err := db.Exec("INSERT INTO data_topic (Name, IsAide, IsBug, IsBoss, IsLore, CategorieID) VALUES (?,?,?,?,?, ?)", top.Name, Aide, Bug, Boss, Lore, top.CategorieID)
	if err != nil {
		return 0, fmt.Errorf("AddTopic: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddTopic: %v", err)
	}
	return id, nil
}
