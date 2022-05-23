package main

//SQL Query that take a topic ID and return it's name
func TakeIDReturnName(ID int64) string {
	var top Data_Topic
	row := db.QueryRow("SELECT Name FROM `data_topic` WHERE ID = ?", ID)
	text := row.Scan(&top.Name)
	if text != nil {
		return text.Error()
	} else {
		return top.Name
	}
}
