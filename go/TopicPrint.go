package main

import "fmt"

// MessagesPrint print message.
func TopicPrint() ([]Data_Topic, error) {

	var topics []Data_Topic

	rows, err := db.Query("SELECT * FROM `data_topic` ORDER BY `data_topic`.`ID` ASC ")
	if err != nil {
		return nil, fmt.Errorf("TopicPrint error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var top Data_Topic
		if err := rows.Scan(&top.ID, &top.Name, &top.IsAide, &top.IsBug, &top.IsBoss, &top.IsLore); err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		top.AddTagsToTopic()
		topics = append(topics, top)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("MessagesPrint error: %v", err)
	}
	return topics, nil
}
