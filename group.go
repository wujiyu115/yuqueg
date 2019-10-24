package yuqueg

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
func (g Group) List(login string) (Groups, error) {
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
		return groups, err
	}
	return groups, nil
}

// Get group
func (g Group) Get(login string) (GroupDetail, error) {
	var gd GroupDetail
	if len(login) == 0 {
		return gd, errors.New("group login or id is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s", login), &gd, EmptyRO)
	if err != nil {
		return gd, err
	}
	return gd, nil
}

// Create group
func (g Group) Create(cg *CreateGroup) (GroupDetail, error) {
	var gd GroupDetail
	if len(cg.Name) == 0 {
		return gd, errors.New("data.name is required")
	}
	if len(cg.Login) == 0 {
		return gd, errors.New("data.login is required")
	}
	_, err := g.client.RequestObj("groups", &gd, &RequestOption{
		Method: "POST",
		Data:   StructToMapStr(cg),
	})
	if err != nil {
		return gd, err
	}
	return gd, nil
}

// Update group
func (g Group) Update(login string, cg *CreateGroup) (GroupDetail, error) {
	var groups GroupDetail

	if len(login) == 0 {
		return groups, errors.New("group login or id is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s", login), &groups, &RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(cg),
	})
	if err != nil {
		return groups, err
	}
	return groups, nil
}

// Delete group
func (g Group) Delete(login string) (GroupDetail, error) {
	var groups GroupDetail
	if len(login) == 0 {
		return groups, errors.New("group login or id is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s", login), &groups, &RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return groups, err
	}
	return groups, nil
}

// ListUsers of group
func (g Group) ListUsers(login string) (GroupUsers, error) {
	var gd GroupUsers
	if len(login) == 0 {
		return gd, errors.New("group login or id is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s/users", login), &gd, EmptyRO)
	if err != nil {
		return gd, err
	}
	return gd, nil
}

// ListUsers of group
func (g Group) AddUser(group string, user string, ga *GroupAddUser) (GroupUserInfo, error) {
	var gd GroupUserInfo

	if len(group) == 0 || len(user) == 0 {
		return gd, errors.New("group and user is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s/users/%s", group, user), &gd, &RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(ga),
	})
	if err != nil {
		return gd, err
	}
	return gd, nil
}

// RemoveUser of group
func (g Group) RemoveUser(group string, user string) (RemoveUserResponse, error) {
	var gd RemoveUserResponse
	if len(group) == 0 || len(user) == 0 {
		return gd, errors.New("group and user is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s/users/%s", group, user), &gd, &RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return gd, err
	}
	return gd, nil
}
