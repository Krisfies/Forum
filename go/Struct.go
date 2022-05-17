package main

type Message struct {
	ID                 int64
	Content            string
	Author             string
	Date               string
	Historique_message []Message
}

type Topic struct {
	ID     int64
	Author string
	Date   string
}

type Data_User struct {
	ID    int64
	Nom   string
	email string
	mdp   string
	admin string
}

type Data_Verif struct {
	Connect_Verif bool
}

type Data_Topic struct {
	Nom string
}
