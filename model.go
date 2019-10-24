package yuqueg

import "time"

var (
	// BaseAPI address
	BaseAPI = "https://www.yuque.com/api/v2/"
	//EmptyRO empty options
	EmptyRO = new(RequestOption)
)

type User struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	SpaceID          int       `json:"space_id"`
	AccountID        int       `json:"account_id"`
	Login            string    `json:"login"`
	Name             string    `json:"name"`
	AvatarURL        string    `json:"avatar_url"`
	LargeAvatarURL   string    `json:"large_avatar_url"`
	MediumAvatarURL  string    `json:"medium_avatar_url"`
	SmallAvatarURL   string    `json:"small_avatar_url"`
	BooksCount       int       `json:"books_count"`
	PublicBooksCount int       `json:"public_books_count"`
	FollowersCount   int       `json:"followers_count"`
	FollowingCount   int       `json:"following_count"`
	Public           int       `json:"public"`
	Description      string    `json:"description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Serializer       string    `json:"_serializer"`
}

//UserInfo JSON structure for the authorized user
type UserInfo struct {
	Data User `json:"data"`
}

//Doc JSON structure for doc
type Doc struct {
	ID     int    `json:"id"`
	Slug   string `json:"slug"`
	Title  string `json:"title"`
	BookID int    `json:"book_id"`
	Book   struct {
		ID               int       `json:"id"`
		Type             string    `json:"type"`
		Slug             string    `json:"slug"`
		Name             string    `json:"name"`
		UserID           int       `json:"user_id"`
		Description      string    `json:"description"`
		CreatorID        int       `json:"creator_id"`
		Public           int       `json:"public"`
		ItemsCount       int       `json:"items_count"`
		LikesCount       int       `json:"likes_count"`
		WatchesCount     int       `json:"watches_count"`
		ContentUpdatedAt time.Time `json:"content_updated_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		CreatedAt        time.Time `json:"created_at"`
		Namespace        string    `json:"namespace"`
		User             struct {
			ID               int         `json:"id"`
			Type             string      `json:"type"`
			Login            string      `json:"login"`
			Name             string      `json:"name"`
			Description      interface{} `json:"description"`
			AvatarURL        string      `json:"avatar_url"`
			LargeAvatarURL   string      `json:"large_avatar_url"`
			MediumAvatarURL  string      `json:"medium_avatar_url"`
			SmallAvatarURL   string      `json:"small_avatar_url"`
			BooksCount       int         `json:"books_count"`
			PublicBooksCount int         `json:"public_books_count"`
			FollowersCount   int         `json:"followers_count"`
			FollowingCount   int         `json:"following_count"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			Serializer       string      `json:"_serializer"`
		} `json:"user"`
		Serializer string `json:"_serializer"`
	} `json:"book"`
	UserID  int `json:"user_id"`
	Creator struct {
		ID               int       `json:"id"`
		Type             string    `json:"type"`
		Login            string    `json:"login"`
		Name             string    `json:"name"`
		Description      string    `json:"description"`
		AvatarURL        string    `json:"avatar_url"`
		LargeAvatarURL   string    `json:"large_avatar_url"`
		MediumAvatarURL  string    `json:"medium_avatar_url"`
		SmallAvatarURL   string    `json:"small_avatar_url"`
		BooksCount       int       `json:"books_count"`
		PublicBooksCount int       `json:"public_books_count"`
		FollowersCount   int       `json:"followers_count"`
		FollowingCount   int       `json:"following_count"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		Serializer       string    `json:"_serializer"`
	} `json:"creator"`
	Format            string      `json:"format"`
	Body              string      `json:"body"`
	BodyDraft         string      `json:"body_draft"`
	BodyHTML          string      `json:"body_html"`
	BodyLake          string      `json:"body_lake"`
	Public            int         `json:"public"`
	Status            int         `json:"status"`
	LikesCount        int         `json:"likes_count"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  time.Time   `json:"content_updated_at"`
	DeletedAt         interface{} `json:"deleted_at"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	FirstPublishedAt  time.Time   `json:"first_published_at"`
	WordCount         int         `json:"word_count"`
	Cover             interface{} `json:"cover"`
	Description       string      `json:"description"`
	CustomDescription interface{} `json:"custom_description"`
	Serializer        string      `json:"_serializer"`
}

// DocDetail JSON structure for a yuque document detail contents
type DocDetail struct {
	Abilities struct {
		Update  bool `json:"update"`
		Destroy bool `json:"destroy"`
	} `json:"abilities"`
	Data Doc `json:"data"`
}

//DocBookDetail of doc
type DocBookDetail struct {
	ID                int         `json:"id"`
	Slug              string      `json:"slug"`
	Title             string      `json:"title"`
	Description       string      `json:"description"`
	UserID            int         `json:"user_id"`
	BookID            int         `json:"book_id"`
	Format            string      `json:"format"`
	Public            int         `json:"public"`
	Status            int         `json:"status"`
	LikesCount        int         `json:"likes_count"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  time.Time   `json:"content_updated_at"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	FirstPublishedAt  time.Time   `json:"first_published_at"`
	DraftVersion      int         `json:"draft_version"`
	LastEditorID      int         `json:"last_editor_id"`
	WordCount         int         `json:"word_count"`
	Cover             interface{} `json:"cover"`
	CustomDescription interface{} `json:"custom_description"`
	LastEditor        struct {
		ID              int       `json:"id"`
		Type            string    `json:"type"`
		Login           string    `json:"login"`
		Name            string    `json:"name"`
		Description     string    `json:"description"`
		AvatarURL       string    `json:"avatar_url"`
		LargeAvatarURL  string    `json:"large_avatar_url"`
		MediumAvatarURL string    `json:"medium_avatar_url"`
		SmallAvatarURL  string    `json:"small_avatar_url"`
		FollowersCount  int       `json:"followers_count"`
		FollowingCount  int       `json:"following_count"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		Serializer      string    `json:"_serializer"`
	} `json:"last_editor"`
	Book       interface{} `json:"book"`
	Serializer string      `json:"_serializer"`
}

// BookDetail  JSON structure for a list of articles in the library
type BookDetail struct {
	Data []DocBookDetail `json:"data"`
}

// Book JSON structure for a list of repositories in the group
type Book struct {
	Data []struct {
		ID               int       `json:"id"`
		Type             string    `json:"type"`
		Slug             string    `json:"slug"`
		Name             string    `json:"name"`
		UserID           int       `json:"user_id"`
		Description      string    `json:"description"`
		CreatorID        int       `json:"creator_id"`
		Public           int       `json:"public"`
		ItemsCount       int       `json:"items_count"`
		LikesCount       int       `json:"likes_count"`
		WatchesCount     int       `json:"watches_count"`
		ContentUpdatedAt time.Time `json:"content_updated_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		CreatedAt        time.Time `json:"created_at"`
		Namespace        string    `json:"namespace"`
		User             User      `json:"user"`
		Serializer       string    `json:"_serializer"`
	} `json:"data"`
}

//Groups JSON structure for  user's groups
type Groups struct {
	Data []struct {
		ID                int       `json:"id"`
		Login             string    `json:"login"`
		Name              string    `json:"name"`
		AvatarURL         string    `json:"avatar_url"`
		LargeAvatarURL    string    `json:"large_avatar_url"`
		MediumAvatarURL   string    `json:"medium_avatar_url"`
		SmallAvatarURL    string    `json:"small_avatar_url"`
		BooksCount        int       `json:"books_count"`
		PublicBooksCount  int       `json:"public_books_count"`
		TopicsCount       int       `json:"topics_count"`
		PublicTopicsCount int       `json:"public_topics_count"`
		MembersCount      int       `json:"members_count"`
		Public            int       `json:"public"`
		Description       string    `json:"description"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Serializer        string    `json:"_serializer"`
	} `json:"data"`
	Abilities struct {
		Read      bool `json:"read"`
		Update    bool `json:"update"`
		Destroy   bool `json:"destroy"`
		GroupUser struct {
			Create  bool `json:"create"`
			Update  bool `json:"update"`
			Destroy bool `json:"destroy"`
		} `json:"group_user"`
		Repo struct {
			Create  bool `json:"create"`
			Update  bool `json:"update"`
			Destroy bool `json:"destroy"`
		} `json:"repo"`
	} `json:"abilities"`
}

//GroupDetail JSON structure for  user's groups
type GroupDetail struct {
	Data struct {
		ID                int       `json:"id"`
		Login             string    `json:"login"`
		Name              string    `json:"name"`
		AvatarURL         string    `json:"avatar_url"`
		LargeAvatarURL    string    `json:"large_avatar_url"`
		MediumAvatarURL   string    `json:"medium_avatar_url"`
		SmallAvatarURL    string    `json:"small_avatar_url"`
		BooksCount        int       `json:"books_count"`
		PublicBooksCount  int       `json:"public_books_count"`
		TopicsCount       int       `json:"topics_count"`
		PublicTopicsCount int       `json:"public_topics_count"`
		MembersCount      int       `json:"members_count"`
		Public            int       `json:"public"`
		OwnerId           int       `json:"owner_id"`
		SpaceId           int       `json:"space_id"`
		Description       string    `json:"description"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Serializer        string    `json:"_serializer"`
	} `json:"data"`
	Abilities struct {
		Read      bool `json:"read"`
		Update    bool `json:"update"`
		Destroy   bool `json:"destroy"`
		GroupUser struct {
			Create  bool `json:"create"`
			Update  bool `json:"update"`
			Destroy bool `json:"destroy"`
		} `json:"group_user"`
		Repo struct {
			Create  bool `json:"create"`
			Update  bool `json:"update"`
			Destroy bool `json:"destroy"`
		} `json:"repo"`
	} `json:"abilities"`
}

type GroupUser struct {
	ID         int       `json:"id"`
	GroupId    int       `json:"group_id"`
	UserId     int       `json:"user_id"`
	Role       int       `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User       User      `json:"user"`
	Serializer string    `json:"_serializer"`
}

type GroupUserInfo struct {
	Data GroupUser `json:"data"`
}

type GroupUsers struct {
	Data []GroupUser `json:"data"`
}

type RemoveUserResponse struct {
	Data struct {
		UserId int `json:"user_id"`
	} `json:"data"`
}

//RepoUser
type RepoUser struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	Login            string    `json:"login"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	AvatarURL        string    `json:"avatar_url"`
	LargeAvatarURL   string    `json:"large_avatar_url"`
	MediumAvatarURL  string    `json:"medium_avatar_url"`
	SmallAvatarURL   string    `json:"small_avatar_url"`
	BooksCount       int       `json:"books_count"`
	PublicBooksCount int       `json:"public_books_count"`
	FollowersCount   int       `json:"followers_count"`
	FollowingCount   int       `json:"following_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Serializer       string    `json:"_serializer"`
}

// UserRepos JSON structure for  user's repos
type UserRepos struct {
	Data []struct {
		ID               int       `json:"id"`
		Type             string    `json:"type"`
		Slug             string    `json:"slug"`
		Name             string    `json:"name"`
		UserID           int       `json:"user_id"`
		Description      string    `json:"description"`
		CreatorID        int       `json:"creator_id"`
		Public           int       `json:"public"`
		ItemsCount       int       `json:"items_count"`
		LikesCount       int       `json:"likes_count"`
		WatchesCount     int       `json:"watches_count"`
		ContentUpdatedAt time.Time `json:"content_updated_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		CreatedAt        time.Time `json:"created_at"`
		Namespace        string    `json:"namespace"`
		User             RepoUser  `json:"user"`
		Serializer       string    `json:"_serializer"`
	} `json:"data"`
}

// UserRepo JSON structure for  create user's repo
type CreateUserRepo struct {
	Data struct {
		ID               int       `json:"id"`
		Type             string    `json:"type"`
		Slug             string    `json:"slug"`
		Name             string    `json:"name"`
		UserID           int       `json:"user_id"`
		Description      string    `json:"description"`
		CreatorID        int       `json:"creator_id"`
		Public           int       `json:"public"`
		ItemsCount       int       `json:"items_count"`
		LikesCount       int       `json:"likes_count"`
		WatchesCount     int       `json:"watches_count"`
		ContentUpdatedAt time.Time `json:"content_updated_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		CreatedAt        time.Time `json:"created_at"`
		Namespace        string    `json:"namespace"`
		User             RepoUser  `json:"user"`
		TocYml           string    `json:"toc_yml"`
		Serializer       string    `json:"_serializer"`
	} `json:"data"`
}

type RepoToc struct {
	Data []struct {
		Title string `json:"title"`
		Slug  string `json:"slug"`
		Depth int    `json:"depth"`
	} `json:"data"`
}
