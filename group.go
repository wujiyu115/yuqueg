package yuque

import (
	"errors"
	"fmt"
)

// Group encapsulate authenticated token
type Group struct {
	client *Client
}

// NewGroup create Doc for external use
func NewGroup(client *Client) *Group {
	return &Group{
		client: client,
	}
}

// List groups
func (g Group) List(login string) (interface{}, error) {
	var (
		url    string
		groups Groups
	)
	if len(login) > 0 {
		url = fmt.Sprintf("users/%s/groups", login)
	} else {
		url = "groups"
	}
	_, err := g.client.RequestObj(url, &groups, EmptyRO)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// Get group
func (g Group) Get(login string) (interface{}, error) {
	if len(login) == 0 {
		return nil, errors.New("group login or id is required")
	}
	var gd GroupDetail
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s", login), &gd, EmptyRO)
	if err != nil {
		return nil, err
	}
	return gd, nil
}

// Create group
func (g Group) Create(cg *CreateGroup) (interface{}, error) {
	if len(cg.Name) == 0 {
		return nil, errors.New("data.name is required")
	}
	if len(cg.Login) == 0 {
		return nil, errors.New("data.login is required")
	}
	var gd GroupDetail
	_, err := g.client.RequestObj("groups", &gd, &RequestOption{
		Method: "POST",
		Data:  StructToMapStr(cg),
	})
	if err != nil {
		return nil, err
	}
	return gd, nil
}

// Update group
func (g Group) Update(login string, cg *CreateGroup) (interface{}, error) {
	if len(login) == 0 {
		return nil, errors.New("group login or id is required")
	}
	var groups GroupDetail
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s", login), &groups, &RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(cg),
	})
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// Delete group
func (g Group) Delete(login string) (interface{}, error) {
	if len(login) == 0 {
		return nil, errors.New("group login or id is required")
	}
	var groups GroupDetail
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s", login), &groups, &RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// ListUsers of group
func (g Group) ListUsers(login string) (interface{}, error) {
	if len(login) == 0 {
		return nil, errors.New("group login or id is required")
	}
	var gd GroupUsers
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s/users", login), &gd, EmptyRO)
	if err != nil {
		return nil, err
	}
	return gd, nil
}

// ListUsers of group
func (g Group) AddUser(group string, user string, ga *GroupAddUser) (interface{}, error) {
	if len(group) == 0 || len(user) == 0 {
		return nil, errors.New("group and user is required")
	}
	var gd GroupUserInfo
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s/users/%s", group, user), &gd, &RequestOption{
		Method: "PUT",
		Data: StructToMapStr(ga),
	})
	if err != nil {
		return nil, err
	}
	return gd, nil
}

// RemoveUser of group
func (g Group) RemoveUser(group string, user string) (interface{}, error) {
	if len(group) == 0 || len(user) == 0 {
		return nil, errors.New("group and user is required")
	}
	var gd RemoveUserResponse
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s/users/%s", group, user), &gd, &RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return nil, err
	}
	return gd, nil
}