package main

import (
	"fmt"
)

func AddUser(usr Data_User) (int64, error) {
	var ID int64
	// Verifie que l'on ne rentre pas une donnée vide lors de l'inscription
	if usr.Name == "" || usr.Password == "" || usr.Email == "" {
		fmt.Println("Erreur le compte ne peux être créée")
		// Crée le profil de l'utilisateur
	} else {
		result, err := db.Exec("INSERT INTO Data_User (Name, Email, Password, Admin, Biography, Date) VALUES (?, ?, ?, ?, ?, ?)", usr.Name, usr.Email, usr.Password, usr.Admin, usr.Biography, usr.Date)
		// Gère les erreurs lors de la création du profil
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
