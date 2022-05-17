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
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	css := http.FileServer(http.Dir("../css/"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	images := http.FileServer(http.Dir("../images/"))
	http.Handle("/images/", http.StripPrefix("/images/", images))

	TemplateMessage := template.Must(template.ParseFiles("../html/TemplateMessage.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tab_messages, err := MessagesPrint()
		if err != nil {
			log.Fatal(err)
		}

		message := Message{
			Content:            r.FormValue("Content"),
			Author:             "Krisfies?",
			Date:               DateMessage(),
			Historique_message: tab_messages,
		}

		if message.Content != " " {
			messID, err := AddMessage(message)
			fmt.Printf("ID of added message: %v\n", messID)
			if err != nil {
				fmt.Println(err)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		TemplateMessage.Execute(w, message)
	})

	http.ListenAndServe(":8010", nil)

}
