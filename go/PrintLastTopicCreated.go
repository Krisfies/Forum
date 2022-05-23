package main

import "fmt"

func PrintLastTopicCreated() ([]Data_Topic, error) {
	var topics []Data_Topic

	rows, err := db.Query("SELECT * FROM `data_topic` WHERE `ID` > (SELECT MAX(`ID`) - 5  FROM `data_topic`) ORDER BY `ID` DESC")
	if err != nil {
		return nil, fmt.Errorf("PrintLastTopicCreated error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var top Data_Topic
		if err := rows.Scan(&top.ID, &top.Name, &top.IsAide, &top.IsBug, &top.IsBoss, &top.IsLore, &top.CategorieID); err != nil {
			return nil, fmt.Errorf("PrintLastTopicCreated error: %v", err)
		}
		top.AddTagsToTopic()
		topics = append(topics, top)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("PrintLastTopicCreated error: %v", err)
	}
	return topics, nil
}
