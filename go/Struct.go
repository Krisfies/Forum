package main

type Data_Message struct {
	ID              int64
	Content         string
	Author          string
	Date            string
	Topic_ID        int64
	Message_History []Data_Message
	Topic_Name      string
}

type Data_User struct {
	ID       int64
	Name     string
	Email    string
	Password string
	Admin    string
}

type Data_Verif struct {
	Connect_Verif bool
}

type Data_Topic struct {
	ID            int64
	Name          string
	Topic_History []Data_Topic
}

type Data_Tags struct {
	ID   int64
	Name string
}
