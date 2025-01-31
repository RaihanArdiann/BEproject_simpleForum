package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/RaihanArdiann/BEproject_simpleForum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postsRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error get user activities from database")
		return err
	}
	if userActivity == nil {
		if !request.IsLiked {
			return errors.New("anda belum pernah like sebelumnya")
		}
		err = s.postsRepo.CreateUserActivity(ctx, model)
	} else {
		err = s.postsRepo.UpdateUserActivity(ctx, model)
	}
	if err != nil {
		log.Error().Err(err).Msg("error create or update user activities")
	}
	return nil
}
