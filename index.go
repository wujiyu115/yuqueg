package yuqueg

// Service encapsulate authenticated token
type Service struct {
	Client *Client
	User   *UserService
	Doc    *DocService
	Repo   *RepoService
	Group  *GroupService
}

// NewService create Client for external use
func NewService(token string) *Service {
	s := new(Service)
	s.Init(token)
	return s
}

func (s *Service) Init(token string) {
	s.Client = NewClient(token)
	s.User = NewUser(s.Client)
	s.Doc = NewDoc(s.Client)
	s.Repo = NewRepo(s.Client)
	s.Group = NewGroup(s.Client)
}
