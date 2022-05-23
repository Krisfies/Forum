package main

import "fmt"

func UserConnection(usr Data_User) {
	rows, err := db.Query("SELECT * FROM `data_user` WHERE `Name` = ?", usr.Name)
	if err != nil {
		fmt.Errorf("TopicPrint error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password, &usr.Admin, &usr.Date, &usr.Biography); err != nil {
			fmt.Errorf("TopicPrint error: %v", err)
		}
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("TopicPrint error: %v", err)
	}
}
