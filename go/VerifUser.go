package main

func VerifUser(UserName string, UserPassword string) bool {
	var alt Data_User
	row := db.QueryRow("SELECT mdp FROM Data_User WHERE Nom = ?", UserName)
	verif := row.Scan(&alt.mdp)
	if verif == nil {
		if UserPassword == alt.mdp {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
