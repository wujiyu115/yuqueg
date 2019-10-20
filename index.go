package main

// Service encapsulate authenticated token
type Service struct {
	client *Client
	user   *UserService
	doc    *Doc
	repo   *Repo
	group  *Group
}

// NewService create Client for external use
func NewService(token string) *Service {
	s := new(Service)
	s.init(token)
	return s
}

func (s *Service) init(token string) {
	s.client = NewClient(token)
	s.user = NewUser(s.client)
	s.doc = NewDoc(s.client)
	s.repo = NewRepo(s.client)
	s.group = NewGroup(s.client)
}
