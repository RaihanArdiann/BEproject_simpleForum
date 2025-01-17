package memberships

import (
	"context"
	"errors"

	"github.com/RaihanArdiann/BEproject_simpleForum/internal/model/memberships"
	"github.com/RaihanArdiann/BEproject_simpleForum/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	log.Info().Msgf("Attempting to login with email: %s", req.Email)
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return "", err
	}
	if user == nil {
		log.Warn().Msg("User not found")
		return "", errors.New("email not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("Password comparison failed")
		return "", errors.New("email or password is invalid")
	}
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", err
	}
	return token, nil
}
