package main

// Structure des messages
type Data_Message struct {
	ID              int64
	Content         string
	Author          string
	Date            string
	Topic_ID        int64
	Message_History []Data_Message
	Topic_Name      string
	Nbr_Like        int64
}

// Structure des utilisateurs
type Data_User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Admin     string
	Biography string
	Date      string
}

// Structure qui vérifie la connection
type Data_Verif struct {
	Connect_Verif bool
}

// Structure des topics
type Data_Topic struct {
	ID            int64
	Name          string
	Topic_History []Data_Topic
	Tags          []string
	IsAide        bool
	IsBug         bool
	IsBoss        bool
	IsLore        bool
	CategorieID   int64
}
