# yuque_go

example:

```golang
package main

import (
	"encoding/json"
	"github.com/wujiyu115/yuqueg"
)

func main() {
	l := yuqueg.L
	yu := yuqueg.NewService("token")
	d, err := yu.User.Get("")

	// d, err := yu.Group.List("u22579")
	// d, err := yu.Group.Get("zutzgg")
	// d, err := yu.Group.Create(&CreateGroup{
	// 	Name: "xiaosz11",
	// 	Login: "xiaosz11",
	// 	Description: "group",
	// })
	// d, err := yu.Group.Update("xiaosz11", &CreateGroup{
	// 	Name: "szxiaocd",
	// 	Login: "szxiaocd",
	// 	Description: "szxiaocd",
	// })
	// d, err := yu.Group.Delete("xiaosz1")
	// d, err := yu.Group.ListUsers("family")
	// d, err := yu.Group.AddUser("zutzgg", "helix", &GroupAddUser{
	// 	Role: 0,
	// })
	// d, err := yu.Group.RemoveUser("zutzgg", "helix")


	// d, err := yu.Repo.List("u22579", "", nil)
	// d, err := yu.Repo.Create("u22579", "", &CreateRepo{
	// 	Name: "xx",
	// 	Slug: "xx",
	// 	Description: "xxx",
	// 	Public: 0,
	// 	Type: "Book",
	// })
	// d, err := yu.Repo.Get("u22579/xx", "Book")
	// d, err := yu.Repo.Update("u22579/xx",  &UpdateRepo{
	// 	Name: "xx1",
	// 	Slug: "xx1",
	// 	Description: "xxx1",
	// 	Public: 0,
	// })
	// d, err := yu.Repo.Delete("u22579/xxx")
	// d, err := yu.Repo.GetToc("u22579/kktbui")

	// d, err := yu.Doc.List("u22579/xcd0mr")
	// d, err := yu.Doc.Get("u22579/xcd0mr", "golang-gai-lan", &DocGet{Raw: 1})
	// d, err := yu.Doc.Create("u22579/xcd0mr", &DocCreate{
	// 	Title: "hexo",
	// 	Slug: "hexo",
	// 	Public: 0,
	// 	Body: "# title \n ## heox is a blog system",
	// })
	// d, err := yu.Doc.Update("u22579/xcd0mr", "2912350", &DocCreate{
	// 	Title: "hexo1",
	// 	Slug: "hexo1",
	// 	Public: 0,
	// 	Body: "# title \n ## heox is a blog system111",
	// })
	// d, err := yu.Doc.Delete("u22579/xcd0mr", "2912350")

	if err != nil {
		l.Info(err)
		return
	}
	jsonString, _ := json.Marshal(d)
	l.Info(string(jsonString))
}

```