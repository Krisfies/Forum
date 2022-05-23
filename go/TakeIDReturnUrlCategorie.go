package main

import "strconv"

func TakeIDReturnUrlCategorie(ID int64) string {
	url := "D" + strconv.Itoa(int(ID))
	return url
}
