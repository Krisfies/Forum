package main

import "fmt"

func Doublon(Nom string) bool {
	var alt Data_User
	row := db.QueryRow("SELECT * FROM Data_User WHERE Name = ?", Nom)
	verif := row.Scan(&alt.ID, &alt.Name, &alt.Email, &alt.Password, &alt.Admin)
	if verif == nil {
		fmt.Println("Ce Pseudo est déja utilisé")
		return true
	} else {
		return false
	}
}
