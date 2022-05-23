package main

func (top *Data_Topic) AddTagsToTopic() {

	if top.IsAide == true {
		top.Tags = append(top.Tags, "#Aide")
	}
	if top.IsBug == true {
		top.Tags = append(top.Tags, "#Bug")
	}
	if top.IsBoss == true {
		top.Tags = append(top.Tags, "#Boss")
	}
	if top.IsLore == true {
		top.Tags = append(top.Tags, "#Lore")
	}
}
