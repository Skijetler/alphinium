package repo

import (
	"context"
	"github.com/Skijetler/alphinium/forum/internal/usecase"
	"github.com/Skijetler/alphinium/pkg/ent"
	entPost "github.com/Skijetler/alphinium/pkg/ent/post"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type forumRepo struct {
	repo *Repo
	log  *logrus.Logger
}

// NewForumRepo .
func NewForumRepo(repo *Repo, logger *logrus.Logger) usecase.ForumRepo {
	rand.Seed(time.Now().UnixNano())

	return &forumRepo{
		repo: repo,
		log:  logger,
	}
}

func (r *forumRepo) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.repo.db.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	if err = fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			r.log.WithContext(ctx).Errorf("Rolling back transaction error: %v", rerr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

// SaveCategory create category and returns it's id and name
func (r *forumRepo) SaveCategory(ctx context.Context, c *usecase.Category) (*usecase.Category, error) {
	var categoryModel *ent.Category

	var err error

	if categoryModel, err = r.repo.db.Category.
		Create().
		SetName(c.Name).
		Save(ctx); err != nil {
		return nil, err
	}

	return &usecase.Category{
		ID:   categoryModel.ID,
		Name: categoryModel.Name,
	}, nil
}

// SaveSubcategory create subcategory and returns it's id, name, desc and category_id
func (r *forumRepo) SaveSubcategory(ctx context.Context, s *usecase.Subcategory) (*usecase.Subcategory, error) {
	var subcategoryModel *ent.Subcategory

	var err error

	if subcategoryModel, err = r.repo.db.Subcategory.
		Create().
		SetName(s.Name).
		SetDescription(s.Description).
		SetCategoryID(s.CategoryID).
		Save(ctx); err != nil {
		return nil, err
	}

	return &usecase.Subcategory{
		ID:          subcategoryModel.ID,
		Name:        subcategoryModel.Name,
		Description: subcategoryModel.Description,
		CategoryID:  subcategoryModel.CategoryID,
	}, nil
}

func addAttachmentsToPost(ctx context.Context, tx *ent.Tx, p *ent.Post, a []usecase.Attachment) ([]usecase.Attachment, error) {
	var postAttachments []usecase.Attachment
	var err error

	for _, attachment := range a {
		var attachmentModel *ent.Attachment

		if attachmentModel, err = tx.Attachment.
			Get(ctx, attachment.ID); err != nil {
			return []usecase.Attachment{}, err
		}

		if p, err = p.
			Update().
			AddAttachments(attachmentModel).
			Save(ctx); err != nil {
			return []usecase.Attachment{}, err
		}

		postAttachments = append(postAttachments, usecase.Attachment{
			ID:   attachmentModel.ID,
			Name: attachmentModel.Name,
			Size: attachmentModel.Size,
			Type: attachmentModel.Type,
			Link: attachmentModel.Link,
		})
	}

	return postAttachments, nil
}

func getAttachmentsFromPost(ctx context.Context, p *ent.Post) ([]usecase.Attachment, error) {
	var postAttachmentsModels []*ent.Attachment
	var postAttachments []usecase.Attachment
	var err error

	if postAttachmentsModels, err = p.
		QueryAttachments().
		All(ctx); err != nil {
		return []usecase.Attachment{}, err
	}

	for _, attachmentModel := range postAttachmentsModels {
		postAttachments = append(postAttachments, usecase.Attachment{
			ID:   attachmentModel.ID,
			Name: attachmentModel.Name,
			Size: attachmentModel.Size,
			Type: attachmentModel.Type,
			Link: attachmentModel.Link,
		})
	}

	return postAttachments, nil
}

// SaveThread create thread and returns it's id, name, desc_post and subcategory_id
func (r *forumRepo) SaveThread(ctx context.Context, t *usecase.Thread) (*usecase.Thread, error) {
	var threadModel *ent.Thread
	var postModel *ent.Post
	var userModel *ent.User
	var userMetadata *ent.UserMetadata
	var postAttachments []usecase.Attachment

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		var err error

		if userModel, err = tx.User.
			Get(ctx, t.Description.User.ID); err != nil {
			return err
		}

		if userMetadata, err = userModel.
			QueryMetadata().
			Only(ctx); err != nil {
			return err
		}

		if threadModel, err = tx.Thread.
			Create().
			SetName(t.Name).
			SetSubcategoryID(t.SubcategoryID).
			Save(ctx); err != nil {
			return err
		}

		if postModel, err = tx.Post.
			Create().
			SetMessage(t.Description.Message).
			SetDate(time.Now().UTC()).
			SetUser(userModel).
			SetThread(threadModel).
			Save(ctx); err != nil {
			return err
		}

		if postAttachments, err = addAttachmentsToPost(ctx, tx, postModel, t.Description.Attachments); err != nil {
			return err
		}

		if threadModel, err = threadModel.
			Update().
			SetDescription(postModel).
			Save(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &usecase.Thread{
		ID:   threadModel.ID,
		Name: threadModel.Name,
		Description: usecase.Post{
			ID: postModel.ID,
			User: usecase.UserEntity{
				ID:        userModel.ID,
				Name:      userModel.Name,
				NameColor: userMetadata.Color,
				Title:     userMetadata.Title,
				Avatar:    "",
				JoinDate:  userModel.RegistrationDate,
			},
			Message:     postModel.Message,
			Date:        postModel.Date,
			Attachments: postAttachments,
			ThreadID:    postModel.ThreadID,
		},
		SubcategoryID: threadModel.SubcategoryID,
	}, nil
}

// SavePost create post and returns it's id, user, message, date, attachments and thread_id
func (r *forumRepo) SavePost(ctx context.Context, p *usecase.Post) (*usecase.Post, error) {
	var postModel *ent.Post
	var userModel *ent.User
	var userMetadata *ent.UserMetadata
	var postAttachments []usecase.Attachment

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		var err error

		if userModel, err = tx.User.
			Get(ctx, p.User.ID); err != nil {
			return err
		}

		if userMetadata, err = userModel.
			QueryMetadata().
			Only(ctx); err != nil {
			return err
		}

		if postModel, err = tx.Post.
			Create().
			SetMessage(p.Message).
			SetDate(time.Now().UTC()).
			SetUser(userModel).
			SetThreadID(p.ThreadID).
			Save(ctx); err != nil {
			return err
		}

		if postAttachments, err = addAttachmentsToPost(ctx, tx, postModel, p.Attachments); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &usecase.Post{
		ID: postModel.ID,
		User: usecase.UserEntity{
			ID:        userModel.ID,
			Name:      userModel.Name,
			NameColor: userMetadata.Color,
			Title:     userMetadata.Title,
			Avatar:    "",
			JoinDate:  userModel.RegistrationDate,
		},
		Message:     postModel.Message,
		Date:        postModel.Date,
		Attachments: postAttachments,
		ThreadID:    postModel.ThreadID,
	}, nil
}

func (r *forumRepo) GetPost(ctx context.Context, postId uint64) (*usecase.Post, error) {
	var postModel *ent.Post
	var userModel *ent.User
	var userMetadata *ent.UserMetadata
	var postAttachments []usecase.Attachment

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		var err error

		if postModel, err = tx.Post.
			Get(ctx, postId); err != nil {
			return err
		}

		if userModel, err = postModel.
			QueryUser().
			Only(ctx); err != nil {
			return err
		}

		if userMetadata, err = userModel.
			QueryMetadata().
			Only(ctx); err != nil {
			return err
		}

		if postAttachments, err = getAttachmentsFromPost(ctx, postModel); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &usecase.Post{
		ID: postModel.ID,
		User: usecase.UserEntity{
			ID:        userModel.ID,
			Name:      userModel.Name,
			NameColor: userMetadata.Color,
			Title:     userMetadata.Title,
			Avatar:    "",
			JoinDate:  userModel.RegistrationDate,
		},
		Message:     postModel.Message,
		Date:        postModel.Date,
		Attachments: postAttachments,
		ThreadID:    postModel.ThreadID,
	}, nil
}

func (r *forumRepo) GetThread(ctx context.Context, threadId uint64) (*usecase.Thread, error) {
	var threadModel *ent.Thread
	var descriptionPost *usecase.Post

	var err error

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		if threadModel, err = tx.Thread.
			Get(ctx, threadId); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if descriptionPost, err = r.GetPost(ctx, threadModel.DescriptionID); err != nil {
		return nil, err
	}

	return &usecase.Thread{
		ID:            threadModel.ID,
		Name:          threadModel.Name,
		Description:   *descriptionPost,
		SubcategoryID: threadModel.SubcategoryID,
	}, nil
}

func (r *forumRepo) GetSubcategory(ctx context.Context, subcategoryId uint64) (*usecase.Subcategory, error) {
	var subcategoryModel *ent.Subcategory
	var subcategoryThreadsModels []*ent.Thread
	var subcategoryThreads []uint64

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		var err error

		if subcategoryModel, err = tx.Subcategory.
			Get(ctx, subcategoryId); err != nil {
			return err
		}

		if subcategoryThreadsModels, err = subcategoryModel.
			QueryThreads().
			All(ctx); err != nil {
			return err
		}

		for _, threadModel := range subcategoryThreadsModels {
			subcategoryThreads = append(subcategoryThreads, threadModel.ID)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &usecase.Subcategory{
		ID:          subcategoryModel.ID,
		Name:        subcategoryModel.Name,
		Description: subcategoryModel.Description,
		CategoryID:  subcategoryModel.CategoryID,
		Threads:     subcategoryThreads,
	}, nil
}

func (r *forumRepo) GetCategory(ctx context.Context, categoryId uint64) (*usecase.Category, error) {
	var categoryModel *ent.Category
	var categorySubcategoriesModels []*ent.Subcategory
	var categorySubcategories []uint64

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		var err error

		if categoryModel, err = tx.Category.
			Get(ctx, categoryId); err != nil {
			return err
		}

		if categorySubcategoriesModels, err = categoryModel.
			QuerySubcategories().
			All(ctx); err != nil {
			return err
		}

		for _, subcategoryModel := range categorySubcategoriesModels {
			categorySubcategories = append(categorySubcategories, subcategoryModel.ID)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &usecase.Category{
		ID:            categoryModel.ID,
		Name:          categoryModel.Name,
		Subcategories: categorySubcategories,
	}, nil
}

func (r *forumRepo) GetLastPosts(ctx context.Context, num int) ([]*usecase.Post, error) {
	var posts []*usecase.Post
	var postsIds []uint64

	var err error

	if postsIds, err = r.repo.db.Post.
		Query().
		Order(ent.Desc(entPost.FieldDate)).
		Limit(num).
		IDs(ctx); err != nil {
		return nil, err
	}

	for _, postId := range postsIds {
		var post *usecase.Post

		if post, err = r.GetPost(ctx, postId); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *forumRepo) GetAllCategories(ctx context.Context) ([]*usecase.Category, error) {
	var categoriesModels []*ent.Category
	var categories []*usecase.Category

	var err error

	if categoriesModels, err = r.repo.db.Category.
		Query().
		All(ctx); err != nil {
		return nil, err
	}

	for _, categoryModel := range categoriesModels {
		categories = append(categories, &usecase.Category{
			ID:   categoryModel.ID,
			Name: categoryModel.Name,
		})
	}

	return categories, nil
}

func (r *forumRepo) DeletePost(ctx context.Context, postId uint64) error {
	var err error

	err = r.repo.db.Post.DeleteOneID(postId).Exec(ctx)

	return err
}

func (r *forumRepo) DeleteThread(ctx context.Context, threadId uint64) error {
	var err error

	err = r.repo.db.Thread.DeleteOneID(threadId).Exec(ctx)

	return err
}

func (r *forumRepo) DeleteSubcategory(ctx context.Context, subcategoryId uint64) error {
	var err error

	err = r.repo.db.Subcategory.DeleteOneID(subcategoryId).Exec(ctx)

	return err
}

func (r *forumRepo) DeleteCategory(ctx context.Context, categoryId uint64) error {
	var err error

	err = r.repo.db.Category.DeleteOneID(categoryId).Exec(ctx)

	return err
}

func (r *forumRepo) GetThreadPostsIds(ctx context.Context, threadId uint64, offset int, postsNum int) ([]uint64, error) {
	var postsIds []uint64

	var err error

	if postsIds, err = r.repo.db.Post.
		Query().
		Where(entPost.ThreadID(threadId)).
		Offset(offset).
		Limit(postsNum).
		IDs(ctx); err != nil {
		return nil, err
	}

	return postsIds, nil
}
