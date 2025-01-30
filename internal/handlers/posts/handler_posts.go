package posts

import (
	"context"

	"github.com/RaihanArdiann/BEproject_simpleForum/internal/middleware"
	"github.com/RaihanArdiann/BEproject_simpleForum/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postsService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
}

type Handler struct {
	*gin.Engine

	postSvc postsService
}

func NewHandler(api *gin.Engine, postsSvc postsService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postsSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)

}
