package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"mine-server-manager/internal/config"
	"mine-server-manager/internal/internalErrors"
	"mine-server-manager/internal/repository"
	"mine-server-manager/internal/services/auth"
	"mine-server-manager/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type MockSQLRepository struct {
	users map[string]repository.User
}

func NewMockSQlRepository() *MockSQLRepository {
	return &MockSQLRepository{
		users: make(map[string]repository.User),
	}
}

func (m *MockSQLRepository) CreateUser(ctx context.Context, args repository.CreateUserParams) (repository.User, error) {
	if _, exists := m.users[args.Email]; exists {
		return repository.User{}, internalErrors.ErrUserAlreadyExists
	}

	user := repository.User{
		ID:       int64(len(m.users) + 1),
		Email:    args.Email,
		Password: args.Password,
		Username: args.Username,
	}

	m.users[args.Email] = user

	return user, nil
}

func (m *MockSQLRepository) GetUserByEmail(ctx context.Context, email string) (repository.User, error) {
	user, ok := m.users[email]
	if !ok {
		return repository.User{}, internalErrors.ErrUserNotFound
	}

	return user, nil
}

func TestAuthHandlers(t *testing.T) {
	mockRepo := NewMockSQlRepository()
	authConfig := &config.AuthConfig{
		Whitelist: []string{"test@example.com"},
		TokenExp:  time.Minute * 15,
		JWTSecret: "test-secret",
	}

	authService := auth.NewAuthService(authConfig, nil, repository.Repository{SqlRepo: mockRepo})
	authHandler := NewAuthHandler(authService)

	router := http.NewServeMux()
	router.HandleFunc("/register", authHandler.RegisterHandler)
	router.HandleFunc("/login", authHandler.LoginHandler)

	server := httptest.NewServer(router)
	defer server.Close()

	t.Run("Successful Registration - success expected", func(t *testing.T) {
		user := models.User{
			Username: "testUser",
			Email:    "test@example.com",
			Password: "password123",
		}

		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", server.URL+"/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request due: %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("exptected status %v; got %v", http.StatusCreated, resp.StatusCode)
		}
	})

	t.Run("Whitelist Check for Registration - failure expected", func(t *testing.T) {
		user := models.User{
			Username: "notWhitelisted",
			Email:    "notWhitelistedUser@example.com",
			Password: "password123",
		}

		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", server.URL+"/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request due: %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusForbidden {
			t.Errorf("exptected status %v; got %v", http.StatusCreated, resp.StatusCode)
		}
	})

	t.Run("Valid login - success expected", func(t *testing.T) {
		user := models.User{
			Username: "testUser",
			Email:    "test@example.com",
			Password: "password123",
		}

		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		mockRepo.users["test@example.com"] = repository.User{
			ID:       int64(len(mockRepo.users) + 1),
			Email:    user.Email,
			Password: string(hashedPwd),
			Username: user.Username,
		}

		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", server.URL+"/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request due: %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("exptected status %v; got %v", http.StatusCreated, resp.StatusCode)
		}

		var successResp models.SuccessResponse
		if err := json.NewDecoder(resp.Body).Decode(&successResp); err != nil {
			t.Fatalf("Failed to decode response due: %v", err)
		}

		if data, ok := successResp.Data.(map[string]any); ok {
			if _, ok := data["token"]; !ok {
				t.Errorf("expected token in response, got %v", data)
			}
		}

	})

	t.Run("Invalid login by password - failure expected", func(t *testing.T) {
		user := models.User{
			Username: "testUser",
			Email:    "test@example.com",
			Password: "passxord123",
		}

		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		mockRepo.users["test@example.com"] = repository.User{
			ID:       int64(len(mockRepo.users) + 1),
			Email:    user.Email,
			Password: string(hashedPwd),
			Username: user.Username,
		}

		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", server.URL+"/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request due: %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("exptected status %v; got %v", http.StatusCreated, resp.StatusCode)
		}
	})

	t.Run("Invalid login by whitelist - failure expected", func(t *testing.T) {
		user := models.User{
			Username: "testUser",
			Email:    "test@rxample.com",
			Password: "passxord123",
		}

		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		mockRepo.users["test@example.com"] = repository.User{
			ID:       int64(len(mockRepo.users) + 1),
			Email:    user.Email,
			Password: string(hashedPwd),
			Username: user.Username,
		}

		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", server.URL+"/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request due: %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusForbidden {
			t.Errorf("exptected status %v; got %v", http.StatusUnauthorized, resp.StatusCode)
		}
	})

	t.Run("Jwt Expiration", func(t *testing.T) {
		shortExpAuthConfig := &config.AuthConfig{
			TokenExp:  time.Millisecond * 1,
			JWTSecret: "test-secret",
		}

		user := models.User{
			Email: "test@rxample.com",
		}

		service := auth.NewAuthService(shortExpAuthConfig, nil, repository.Repository{SqlRepo: mockRepo})

		token, err := service.GenerateToken(user.Email)
		if err != nil {
			t.Fatalf("unable to generate token due: %v", err)
		}

		time.Sleep(time.Millisecond * 2)

		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		rr := httptest.NewRecorder()

		protectedHandler := service.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		protectedHandler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusUnauthorized {
			t.Errorf("handler returned wrongstatus code: got %v want %v", status, http.StatusUnauthorized)
		}
	})
}
