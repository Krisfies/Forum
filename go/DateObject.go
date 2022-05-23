package main

import (
	"fmt"
	"time"
)

func DateObject() string {
	var Date string
	DateNow := time.Now()
	Test := fmt.Sprintln(DateNow)
Loop:
	for _, letter := range Test {
		if letter == ' ' {
			break Loop
		}
		Date += string(letter)
	}
	return (Date)
}
