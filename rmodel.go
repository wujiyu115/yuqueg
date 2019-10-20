package main

//CreateGroup struct for create group
type CreateGroup struct {
	Name string				`name`
	Login string			`login`
	Description string		`description`
}

//CreateRepo struct for create repo
type CreateRepo struct {
	Name string				`name`
	Slug string				`slug`
	Description string		`description`
	// 0 私密, 1 内网公开, 2 全网公开
	Public int				`public`
	// ‘Book’ 文库, ‘Design’ 画板, 请注意大小写
	Type string				`type`
}

//UpdateRepo struct for update repo
type UpdateRepo struct {
	Name string				`name`
	Slug string				`slug`
	Description string		`description`
	// 0 私密, 1 内网公开, 2 全网公开
	Public int				`public`
	Toc string 				`toc`
}

//GroupAddUser struct for update repo
type GroupAddUser struct {
	Role 	int		`role`
}

//DocGet struct for get doc
type DocGet struct {
	Raw 	int		`raw`
}
//DocCreate struct for create doc
type DocCreate struct {
	Title 	string		`title`
	Slug 	string		`slug`
	Public	int			`public`
	// markdown or lake, default is markdown
	Format  string		`format`
	//max 5Mb
	Body	string		`body`
}