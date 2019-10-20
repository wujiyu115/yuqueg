package yuque

import (
	"errors"
	"fmt"
)

// Repo encapsulate authenticated token
type Repo struct {
	client *Client
}

// NewRepo create User for external use
func NewRepo(client *Client) *Repo {
	return &Repo{
		client: client,
	}
}

// List url
func (r Repo) List(user string, group string, data map[string]string) (interface{}, error) {
	if len(user) == 0 && len(group) == 0 {
		return nil, errors.New("user or group is required")
	}
	var (
		url string
		u   UserRepos
	)
	if len(user) > 0 {
		url = fmt.Sprintf("users/%s/repos", user)
	} else {
		url = fmt.Sprintf("groups/%s/repos", group)
	}
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Data: data,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Create repo
func (r Repo) Create(user string, group string, cr *CreateRepo) (interface{}, error) {
	if len(user) == 0 && len(group) == 0 {
		return nil, errors.New("user or group is required")
	}
	var (
		url string
		u   CreateUserRepo
	)
	if len(user) > 0 {
		url = fmt.Sprintf("users/%s/repos", user)
	} else {
		url = fmt.Sprintf("groups/%s/repos", group)
	}
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Method: "POST",
		Data: StructToMapStr(cr),
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Get repo
func (r Repo) Get(namespace string, t string) (interface{}, error) {
	if len(namespace) == 0 && len(t) == 0 {
		return nil, errors.New("namespace or type is required")
	}
	var u CreateUserRepo
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Data: map[string]string{"type": t},
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Update repo
func (r Repo) Update(namespace string, cr *UpdateRepo) (interface{}, error) {
	if len(namespace) == 0 {
		return nil, errors.New("namespace is required")
	}
	var u CreateUserRepo
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Method: "PUT",
		Data: StructToMapStr(cr),
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Delete repo
func (r Repo) Delete(namespace string) (interface{}, error) {
	if len(namespace) == 0 {
		return nil, errors.New("namespace is required")
	}
	var u CreateUserRepo
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

//GetToc of repo
func (r Repo) GetToc(namespace string) (interface{}, error) {
	if len(namespace) == 0 {
		return nil, errors.New("namespace is required")
	}
	var u RepoToc
	url := fmt.Sprintf("repos/%s/toc/", namespace)
	_, err := r.client.RequestObj(url, &u, EmptyRO)
	if err != nil {
		return nil, err
	}
	return u, nil
}