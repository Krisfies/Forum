package main

import (
	"fmt"
)

func AddData_User(dat Data_User) (int64, error) {
	if dat.Name == "" || dat.Password == "" || dat.Email == "" {
		fmt.Println("Erreur le compte ne peux être créée")
	} else {
		result, err := db.Exec("INSERT INTO Data_User (Name, Email, Password, Admin) VALUES (?, ?, ?, ?)", dat.Name, dat.Email, dat.Password, dat.Admin)
		if err != nil {
			return 0, fmt.Errorf("addData_User: %v", err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("addData_User: %v", err)
		}
		//defer result.Close()
		return id, nil
	}
	return 0, nil
}
