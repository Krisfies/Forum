package main

func VerifUser(UserName string, UserPassword string) bool {
	var alt Data_User
	// On selectionne un utilisateur par rapport à son pseudo unique
	row := db.QueryRow("SELECT Password FROM Data_User WHERE Name = ?", UserName)
	// On recupère son mot de passe
	verif := row.Scan(&alt.Password)
	// On verifie que le mot de passe entré correpond bien au mot de passe de la base de donnée
	if verif == nil {
		if UserPassword == alt.Password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
