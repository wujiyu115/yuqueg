package yuqueg

// Service encapsulate authenticated token
type Service struct {
	Client *Client
	User   *UserService
	Doc    *Doc
	Repo   *Repo
	Group  *Group
}

// NewService create Client for external use
func NewService(token string) *Service {
	s := new(Service)
	s.init(token)
	return s
}

func (s *Service) init(token string) {
	s.Client = NewClient(token)
	s.User = NewUser(s.Client)
	s.Doc = NewDoc(s.Client)
	s.Repo = NewRepo(s.Client)
	s.Group = NewGroup(s.Client)
}
