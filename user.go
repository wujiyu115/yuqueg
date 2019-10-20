package yuque

import (
	"fmt"
)

// User encapsulate authenticated token
type UserService struct {
	client *Client
}

// NewUser create User for external use
func NewUser(client *Client) *UserService {
	return &UserService{
		client: client,
	}
}

// Request url
func (c UserService) Get(login string) (interface{}, error) {
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
		return nil, err
	}
	return user, nil
}
