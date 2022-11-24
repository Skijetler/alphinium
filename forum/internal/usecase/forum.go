package usecase

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"strconv"
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
	ID        uint64
	Name      string
	NameColor string
	Title     string
	Avatar    string
	JoinDate  time.Time
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
	GetPost(context.Context, uint64) (*Post, error)
	GetThread(context.Context, uint64) (*Thread, error)
	GetSubcategory(context.Context, uint64) (*Subcategory, error)
	GetCategory(context.Context, uint64) (*Category, error)
	GetLastPosts(context.Context, int) ([]*Post, error)
	GetAllCategories(context.Context) ([]*Category, error)
	DeletePost(context.Context, uint64) error
	DeleteThread(context.Context, uint64) error
	DeleteSubcategory(context.Context, uint64) error
	DeleteCategory(context.Context, uint64) error
	GetThreadPostsIds(context.Context, uint64, int, int) ([]uint64, error)
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

func (uc *ForumUsecase) extractUserID(ctx context.Context) (uint64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, internalErr(errors.New("metadata not found"))
	}
	mdValue := md.Get("user_id")
	if len(mdValue) == 0 {
		return 0, internalErr(errors.New("empty metadata"))
	}
	userId, err := strconv.ParseUint(mdValue[0], 10, 64)
	if err != nil {
		return 0, internalErr(errors.New("invalid metadata"))
	}
	return userId, nil
}

func (uc *ForumUsecase) GetAllCategories(ctx context.Context) ([]*Category, error) {
	categories, err := uc.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (uc *ForumUsecase) GetCategory(ctx context.Context, categoryId uint64) (*Category, error) {
	category, err := uc.repo.GetCategory(ctx, categoryId)
	if err != nil {
		return nil, ErrCategoryNotFound
	}

	return category, nil
}

func (uc *ForumUsecase) SaveCategory(ctx context.Context, c *Category) (uint64, error) {
	saved, err := uc.repo.SaveCategory(ctx, c)
	if err != nil {
		return 0, err
	}

	return saved.ID, nil
}
