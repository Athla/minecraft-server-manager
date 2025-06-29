package auth

import (
	"crypto/subtle"
	"log/slog"
	"mine-server-manager/internal/config"
	"mine-server-manager/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	cfg    *config.AuthConfig
	logger *slog.Logger

	db repository.Repository
}

func NewAuthService(cfg *config.AuthConfig, logger *slog.Logger) *AuthService {
	return &AuthService{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *AuthService) IsWhitelisted(userEmail string) bool {
	for _, u := range s.cfg.Whitelist {
		if subtle.ConstantTimeCompare([]byte(u), []byte(userEmail)) == 1 {
			return true
		}
	}

	return false
}

func (s *AuthService) GenerateToken(userEmail string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userEmail,
		"exp": time.Now().Add(s.cfg.TokenExp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.cfg.JWTSecret))
}

func (s *AuthService) ValidatePwd(userEmail, inputPwd string) bool {
	hashedPwd, err := s.db.SqlRepo.RetrieveHashedPwd(userEmail)
	// short hand if flow

	err = bcrypt.CompareHashAndPassword(hashedPwd, []byte(inputPwd))
	return err == nil
}
