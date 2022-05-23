package main

func Citation(ID int64) string {
	var alt Data_Message
	// Selectionne le message qui correspond à l'ID en argument
	row := db.QueryRow("SELECT Content FROM data_message WHERE ID = ?", ID)
	// Contient le message
	text := row.Scan(&alt.ID, &alt.Content)
	if text == nil {
		return alt.Content
	} else {
		return alt.Content
	}
}
