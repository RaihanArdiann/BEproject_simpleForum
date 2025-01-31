package posts

import (
	"context"

	"github.com/RaihanArdiann/BEproject_simpleForum/internal/configs"
	"github.com/RaihanArdiann/BEproject_simpleForum/internal/model/posts"
)

type postsRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
}

type service struct {
	cfg       *configs.Config
	postsRepo postsRepository
}

func NewService(cfg *configs.Config, postsRepo postsRepository) *service {
	return &service{
		cfg:       cfg,
		postsRepo: postsRepo,
	}
}
