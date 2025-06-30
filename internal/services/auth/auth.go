package auth

import (
	"context"
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

	db *repository.Repository
}

func NewAuthService(cfg *config.AuthConfig, logger *slog.Logger, db *repository.Repository) *AuthService {
	return &AuthService{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (s *AuthService) Login(ctx context.Context, email, inputPwd string) (string, error) {
	currUsr, err := s.db.SqlRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(currUsr.Password), []byte(inputPwd))
	if err != nil {
		return "", err
	}

	token, err := s.GenerateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Logout(tokenString string) error {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(s.cfg.JWTSecret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return s.db.CacheRepository.Add(tokenString, "invalidated", 0)
	}

	expiration := claims.ExpiresAt.Time
	ttl := time.Until(expiration)

	return s.db.CacheRepository.Add(tokenString, "invalidated", ttl)
}

func (s *AuthService) CreateUser(ctx context.Context, userName, userEmail, pwd string) (*repository.User, error) {
	const HASH_COST = 10
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), HASH_COST)
	if err != nil {
		return nil, err
	}

	params := repository.CreateUserParams{
		Username: userName,
		Email:    userEmail,
		Password: string(hashedPwd),
	}

	createdUser, err := s.db.SqlRepo.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
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

func (s *AuthService) ValidatePwd(ctx context.Context, userEmail, inputPwd string) bool {
	currUsr, err := s.db.SqlRepo.GetUserByEmail(ctx, userEmail)
	// short hand if flow

	err = bcrypt.CompareHashAndPassword([]byte(currUsr.Password), []byte(inputPwd))
	return err == nil
}
