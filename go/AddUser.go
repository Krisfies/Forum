package main

import (
	"fmt"
)

func AddData_User(dat Data_User) (int64, error) {
	if dat.Nom == "" || dat.mdp == "" || dat.email == "" {
		fmt.Println("Erreur le compte ne peux être créée")
	} else {
		result, err := db.Exec("INSERT INTO Data_User (nom, email, mdp, admin) VALUES (?, ?, ?, ?)", dat.Nom, dat.email, dat.mdp, dat.admin)
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
