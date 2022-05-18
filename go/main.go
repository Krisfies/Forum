package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "data-access",
		AllowNativePasswords: true,
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	}

	fmt.Println("Connected!")
	User_Profil := Data_User{
		Name:     "Random",
		Password: "*******",
		Admin:    "false",
	}

	index := template.Must(template.ParseFiles("../html/inscrire.html"))
	http.HandleFunc("/inscrire", func(w http.ResponseWriter, r *http.Request) {
		UserInput := r.FormValue("UserInput")
		UserInput2 := r.FormValue("UserInput2")
		UserInput3 := r.FormValue("UserInput3")
		NewUser, err := AddData_User(Data_User{
			ID:       000000,
			Name:     UserInput,
			Email:    UserInput2,
			Password: UserInput3,
			Admin:    "false",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(NewUser)
		index.Execute(w, NewUser)
	})

	index2 := template.Must(template.ParseFiles("../html/connecter.html"))
	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		User := Data_User{
			Name:     r.FormValue("UserInput"),
			Password: r.FormValue("UserInput2"),
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(User.Name)
		fmt.Println(User.Password)
		Verif := Data_Verif{
			Connect_Verif: VerifUser(User.Name, User.Password),
		}
		Verif.Connect_Verif = VerifUser(User.Name, User.Password)
		if Verif.Connect_Verif {
			User_Profil.Name = User.Name
			http.Redirect(w, r, "/principal", http.StatusSeeOther)
		}
		index2.Execute(w, User)
	})

	index3 := template.Must(template.ParseFiles("../html/principal.html"))
	http.HandleFunc("/principal", func(w http.ResponseWriter, r *http.Request) {
		Topics := Data_Topic{
			Name: "Topic1",
		}
		fmt.Println(Topics)
		index3.Execute(w, Topics)
	})

	D1 := template.Must(template.ParseFiles("../html/Souls1.html"))
	http.HandleFunc("/D1", func(w http.ResponseWriter, r *http.Request) {

		tab_topic, err := TopicPrint()
		if err != nil {
			fmt.Println(err)
		}

		Topic := Data_Topic{
			Topic_History: tab_topic,
			Name:          "HHEYYYYYYYYYYYYYYBOYIMFROMTEXAS",
		}

		D1.Execute(w, Topic)

	})

	D2 := template.Must(template.ParseFiles("../html/Souls2.html"))
	http.HandleFunc("/D2", func(w http.ResponseWriter, r *http.Request) {

		D2.Execute(w, nil)

	})

	D3 := template.Must(template.ParseFiles("../html/Souls3.html"))
	http.HandleFunc("/D3", func(w http.ResponseWriter, r *http.Request) {

		D3.Execute(w, nil)

	})

	TemplateMessage := template.Must(template.ParseFiles("../html/TemplateMessage.html"))
	http.HandleFunc("/topic", func(w http.ResponseWriter, r *http.Request) {

		tab_messages, err := MessagesPrint()
		if err != nil {
			fmt.Println(err)
		}

		message := Data_Message{
			Content:         r.FormValue("Content"),
			Author:          User_Profil.Name,
			Date:            DateMessage(),
			Message_History: tab_messages,
		}

		if message.Content != "" {
			messID, err := AddMessage(message)
			fmt.Printf("ID of added message: %v\n", messID)
			if err != nil {
				fmt.Println(err)
				return
			}
			http.Redirect(w, r, "/topic", http.StatusSeeOther)
		}

		TemplateMessage.Execute(w, message)
	})

	css := http.FileServer(http.Dir("../css/"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	images := http.FileServer(http.Dir("../images/"))
	http.Handle("/images/", http.StripPrefix("/images/", images))
	http.ListenAndServe(":8010", nil)

}
