package main

import (
	"encoding/json"
)

func main() {
	yu := NewService("xxx")
	d, err := yu.user.Get("")

	// d, err := yu.group.List("u22579")
	// d, err := yu.group.Get("zutzgg")
	// d, err := yu.group.Create(&CreateGroup{
	// 	Name: "xiaosz11",
	// 	Login: "xiaosz11",
	// 	Description: "group",
	// })
	// d, err := yu.group.Update("xiaosz11", &CreateGroup{
	// 	Name: "szxiaocd",
	// 	Login: "szxiaocd",
	// 	Description: "szxiaocd",
	// })
	// d, err := yu.group.Delete("xiaosz1")
	// d, err := yu.group.ListUsers("family")
	// d, err := yu.group.AddUser("zutzgg", "helix", &GroupAddUser{
	// 	Role: 0,
	// })
	// d, err := yu.group.RemoveUser("zutzgg", "helix")


	// d, err := yu.repo.List("u22579", "", nil)
	// d, err := yu.repo.Create("u22579", "", &CreateRepo{
	// 	Name: "xx",
	// 	Slug: "xx",
	// 	Description: "xxx",
	// 	Public: 0,
	// 	Type: "Book",
	// })
	// d, err := yu.repo.Get("u22579/xx", "Book")
	// d, err := yu.repo.Update("u22579/xx",  &UpdateRepo{
	// 	Name: "xx1",
	// 	Slug: "xx1",
	// 	Description: "xxx1",
	// 	Public: 0,
	// })
	// d, err := yu.repo.Delete("u22579/xxx")
	// d, err := yu.repo.GetToc("u22579/kktbui")

	// d, err := yu.doc.List("u22579/xcd0mr")
	// d, err := yu.doc.Get("u22579/xcd0mr", "golang-gai-lan", &DocGet{Raw: 1})
	// d, err := yu.doc.Create("u22579/xcd0mr", &DocCreate{
	// 	Title: "hexo",
	// 	Slug: "hexo",
	// 	Public: 0,
	// 	Body: "# title \n ## heox is a blog system",
	// })
	// d, err := yu.doc.Update("u22579/xcd0mr", "2912350", &DocCreate{
	// 	Title: "hexo1",
	// 	Slug: "hexo1",
	// 	Public: 0,
	// 	Body: "# title \n ## heox is a blog system111",
	// })
	// d, err := yu.doc.Delete("u22579/xcd0mr", "2912350")

	if err != nil {
		l.Info(err)
		return
	}
	jsonString, _ := json.Marshal(d)
	l.Info(string(jsonString))
}
