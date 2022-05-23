package main

import (
	"fmt"
)

func AddUser(usr Data_User) (int64, error) {
	var ID int64
	if usr.Name == "" || usr.Password == "" || usr.Email == "" {
		fmt.Println("Erreur le compte ne peux être créée")
	} else {
		result, err := db.Exec("INSERT INTO Data_User (Name, Email, Password, Admin, Biography, Date) VALUES (?, ?, ?, ?, ?, ?)", usr.Name, usr.Email, usr.Password, usr.Admin, usr.Biography, usr.Date)
		if err != nil {
			return 0, fmt.Errorf("AddData_User: %v", err)
		}
		id, err := result.LastInsertId()
		ID = id
		if err != nil {
			return 0, fmt.Errorf("AddData_User: %v", err)
		}
		return id, nil
	}
	return ID, nil
}
