package main

import (
	"fmt"
	"strconv"
)

func TakeUrlReturnIDCategorie(url string) int64 {
	var IDstr string
	var Idint int64
	for _, letter := range url {
		if letter != '/' && letter != 'D' {
			IDstr += string(letter)
		}
	}
	Idint, err := strconv.ParseInt(IDstr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return Idint
}
