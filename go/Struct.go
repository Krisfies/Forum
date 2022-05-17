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
