package main

import (
	"fmt"
	"time"
)

func DateMessage() string {
	var Date string
	DateMessage := time.Now()
	Test := fmt.Sprintln(DateMessage)
Loop:
	for _, letter := range Test {
		if letter == ' ' {
			break Loop
		}
		Date += string(letter)
	}
	return (Date)
}
