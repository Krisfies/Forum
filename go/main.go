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
		Passwd:               "password",
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

	Registration := template.Must(template.ParseFiles("../html/Registration.html"))
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
		Registration.Execute(w, NewUser)
	})

	Connection := template.Must(template.ParseFiles("../html/Connection.html"))
	http.HandleFunc("/connecter", func(w http.ResponseWriter, r *http.Request) {
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
		Connection.Execute(w, User)
	})

	MainPage := template.Must(template.ParseFiles("../html/MainPage.html"))
	http.HandleFunc("/principal", func(w http.ResponseWriter, r *http.Request) {

		MainPage.Execute(w, User_Profil)
	})

	Profil := template.Must(template.ParseFiles("../html/Profil.html"))
	http.HandleFunc("/profil", func(w http.ResponseWriter, r *http.Request) {

		Profil.Execute(w, User_Profil)
	})

	D1 := template.Must(template.ParseFiles("../html/Souls1.html"))
	http.HandleFunc("/D1", func(w http.ResponseWriter, r *http.Request) {

		tab_topic, err := TopicPrint()
		if err != nil {
			fmt.Println(err)
		}

		Topic := Data_Topic{
			Topic_History: tab_topic,
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

	CréerTopic := template.Must(template.ParseFiles("../html/TopicCreation.html"))
	http.HandleFunc("/créertopic", func(w http.ResponseWriter, r *http.Request) {

		LeTopic := Data_Topic{
			Name: r.FormValue("Name"),
			Tags: r.Form["Tags"],
		}
		fmt.Println(LeTopic.Tags)

		if LeTopic.Name != "" {
			topID, err := AddTopic(LeTopic)
			fmt.Printf("ID of added topic: %v\n", topID)
			if err != nil {
				fmt.Println(err)
				return
			}

			LeMessage := Data_Message{
				Content:    r.FormValue("Content"),
				Author:     User_Profil.Name,
				Date:       DateMessage(),
				Topic_ID:   topID,
				Topic_Name: TakeIDReturnName(topID),
				Nbr_Like:   0,
			}

			if LeMessage.Content != "" {
				messID, err := AddMessage(LeMessage)
				fmt.Printf("ID of added message: %v\n", messID)
				if err != nil {
					fmt.Println(err)
					return
				}
				http.Redirect(w, r, "/D1", http.StatusSeeOther)
			} else {
				//Erreur si message vide
			}
		} else {
			//Erreur si nom de topic vide
		}

		CréerTopic.Execute(w, LeTopic)

	})

	TemplateMessage := template.Must(template.ParseFiles("../html/TemplateMessage.html"))
	http.HandleFunc("/topic/", func(w http.ResponseWriter, r *http.Request) {

		LeTopicId := TakeUrlReturnID(r.URL.String())

		message := Data_Message{
			Content:    r.FormValue("Content"),
			Author:     User_Profil.Name,
			Date:       DateMessage(),
			Topic_ID:   LeTopicId,
			Topic_Name: TakeIDReturnName(LeTopicId),
		}

		tab_messages, err := MessagesPrint(message.Topic_ID)
		if err != nil {
			fmt.Println(err)
		}

		message.Message_History = tab_messages

		// TopicName, err := SelectTopicName(14)
		// if err != nil {
		// 	fmt.Println(err)
		// 	message.Topic_Name = "Error: name couldn't be found"
		// } else {
		// 	message.Topic_Name = TopicName
		// }

		if message.Content != "" {
			messID, err := AddMessage(message)
			fmt.Printf("ID of added message: %v\n", messID)
			if err != nil {
				fmt.Println(err)
				return
			}
			url := r.URL.String()
			http.Redirect(w, r, url, http.StatusSeeOther)
		}

		TemplateMessage.Execute(w, message)
	})

	css := http.FileServer(http.Dir("../css/"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	images := http.FileServer(http.Dir("../images/"))
	http.Handle("/images/", http.StripPrefix("/images/", images))

	js := http.FileServer(http.Dir("../js/"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	http.ListenAndServe(":8010", nil)

}
