package main

func IsTagInTopic(top Data_Topic) (int, int, int, int) {
	var Aide int = 0
	var Bug int = 0
	var Boss int = 0
	var Lore int = 0

	for _, tag := range top.Tags {
		if tag == "#Aide" {
			Aide = 1
		} else if tag == "#Bug" {
			Bug = 1
		} else if tag == "#Boss" {
			Boss = 1
		} else if tag == "#Lore" {
			Lore = 1
		}
	}
	return Aide, Bug, Boss, Lore
}
