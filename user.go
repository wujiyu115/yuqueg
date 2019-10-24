package yuqueg

import (
	"fmt"
)

// UserService encapsulate authenticated token
type UserService struct {
	client *Client
}

// NewUser create User for external use
func NewUser(client *Client) *UserService {
	return &UserService{
		client: client,
	}
}

// Get user
func (c UserService) Get(login string) (UserInfo, error) {
	var (
		url  string
		user UserInfo
	)
	if len(login) > 0 {
		url = fmt.Sprintf("users/%s", login)
	} else {
		url = "user"
	}
	_, err := c.client.RequestObj(url, &user, EmptyRO)
	if err != nil {
		return user, err
	}
	return user, nil
}
