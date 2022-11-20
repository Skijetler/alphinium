package usecase

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Category struct {
	ID            uint64
	Name          string
	Subcategories []uint64
}

type Subcategory struct {
	ID          uint64
	Name        string
	Description string
	CategoryID  uint64
	Threads     []uint64
}

type Thread struct {
	ID            uint64
	Name          string
	Description   Post
	SubcategoryID uint64
}

type UserEntity struct {
	ID       uint64
	Name     string
	Title    string
	Avatar   string
	JoinDate time.Time
}

type Post struct {
	ID          uint64
	User        UserEntity
	Message     string
	Date        time.Time
	Attachments []Attachment
	ThreadID    uint64
}

type Attachment struct {
	ID   uint64
	Name string
	Size string
	Type string
	Link string
}

// ForumRepo is a Forum repo.
type ForumRepo interface {
	SaveCategory(context.Context, *Category) (*Category, error)
	SaveSubcategory(context.Context, *Subcategory) (*Subcategory, error)
	SaveThread(context.Context, *Thread) (*Thread, error)
	SavePost(context.Context, *Post) (*Post, error)
}

// ForumUsecase is a Forum usecase.
type ForumUsecase struct {
	repo ForumRepo
	log  *logrus.Logger
}

// NewForumUsecase new a Forum usecase.
func NewForumUsecase(repo ForumRepo, logger *logrus.Logger) *ForumUsecase {
	return &ForumUsecase{repo: repo, log: logger}
}
