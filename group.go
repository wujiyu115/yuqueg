package yuqueg

import (
	"errors"
	"fmt"
)

// GroupService encapsulate authenticated token
type GroupService struct {
	client *Client
}

// NewGroup create Doc for external use
func NewGroup(client *Client) *GroupService {
	return &GroupService{
		client: client,
	}
}

// List groups
func (g GroupService) List(login string) (Groups, error) {
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
func (g GroupService) Get(login string) (GroupDetail, error) {
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
func (g GroupService) Create(cg *CreateGroup) (GroupDetail, error) {
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
func (g GroupService) Update(login string, cg *CreateGroup) (GroupDetail, error) {
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
func (g GroupService) Delete(login string) (GroupDetail, error) {
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
func (g GroupService) ListUsers(login string) (GroupUsers, error) {
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
func (g GroupService) AddUser(group string, user string, ga *GroupAddUser) (GroupUserInfo, error) {
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
func (g GroupService) RemoveUser(group string, user string) (RemoveUserResponse, error) {
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
