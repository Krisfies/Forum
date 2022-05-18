package main

func Citation(ID int64) string {
	var alt Message
	row := db.QueryRow("SELECT Content FROM data_message WHERE ID = ?", ID)
	text := row.Scan(&alt.ID, &alt.Content)
	if text == nil {
		return alt.Content
	} else {
		return alt.Content
	}
}
