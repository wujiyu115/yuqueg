package yuqueg

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
func (r Repo) List(user string, group string, data map[string]string) (UserRepos, error) {
	var (
		url string
		u   UserRepos
	)
	if len(user) == 0 && len(group) == 0 {
		return u, errors.New("user or group is required")
	}
	if len(user) > 0 {
		url = fmt.Sprintf("users/%s/repos", user)
	} else {
		url = fmt.Sprintf("groups/%s/repos", group)
	}
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Data: data,
	})
	if err != nil {
		return u, err
	}
	return u, nil
}

// Create repo
func (r Repo) Create(user string, group string, cr *CreateRepo) (CreateUserRepo, error) {
	var (
		url string
		u   CreateUserRepo
	)
	if len(user) == 0 && len(group) == 0 {
		return u, errors.New("user or group is required")
	}
	if len(user) > 0 {
		url = fmt.Sprintf("users/%s/repos", user)
	} else {
		url = fmt.Sprintf("groups/%s/repos", group)
	}
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Method: "POST",
		Data:   StructToMapStr(cr),
	})
	if err != nil {
		return u, err
	}
	return u, nil
}

// Get repo
func (r Repo) Get(namespace string, t string) (CreateUserRepo, error) {
	var u CreateUserRepo

	if len(namespace) == 0 && len(t) == 0 {
		return u, errors.New("namespace or type is required")
	}
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Data: map[string]string{"type": t},
	})
	if err != nil {
		return u, err
	}
	return u, nil
}

//Update repo
func (r Repo) Update(namespace string, cr *UpdateRepo) (CreateUserRepo, error) {
	var u CreateUserRepo

	if len(namespace) == 0 {
		return u, errors.New("namespace is required")
	}
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(cr),
	})
	if err != nil {
		return u, err
	}
	return u, nil
}

//Delete repo
func (r Repo) Delete(namespace string) (CreateUserRepo, error) {
	var u CreateUserRepo
	if len(namespace) == 0 {
		return u, errors.New("namespace is required")
	}
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return u, err
	}
	return u, nil
}

//GetToc of repo
func (r Repo) GetToc(namespace string) (RepoToc, error) {
	var u RepoToc
	if len(namespace) == 0 {
		return u, errors.New("namespace is required")
	}
	url := fmt.Sprintf("repos/%s/toc/", namespace)
	_, err := r.client.RequestObj(url, &u, EmptyRO)
	if err != nil {
		return u, err
	}
	return u, nil
}
