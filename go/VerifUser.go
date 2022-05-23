package main

func VerifUser(UserName string, UserPassword string) bool {
	var alt Data_User
	row := db.QueryRow("SELECT Password FROM Data_User WHERE Name = ?", UserName)
	verif := row.Scan(&alt.Password)
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
