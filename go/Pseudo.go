package main

import "fmt"

func Doublon(Nom string) bool {
	var alt Data_User
	// On verifie que le nom n'est pas présent dans la base de donnée
	row := db.QueryRow("SELECT * FROM Data_User WHERE Name = ?", Nom)
	verif := row.Scan(&alt.ID, &alt.Name, &alt.Email, &alt.Password, &alt.Admin)
	// Si la verification renvoie nil alors le pseudo est dans la base de donnée
	if verif == nil {
		fmt.Println("Ce Pseudo est déja utilisé")
		return true
	} else {
		return false
	}
}
