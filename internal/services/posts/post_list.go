package posts

import (
	"context"

	"github.com/RaihanArdiann/BEproject_simpleForum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostresponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)

	response, err := s.postsRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error get all post from database")
		return response, err
	}
	return response, nil
}
