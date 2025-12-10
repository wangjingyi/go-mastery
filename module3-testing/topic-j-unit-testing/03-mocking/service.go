// Package service demonstrates mocking with interfaces
package service

import "context"

// UserRepository defines the interface for user data access
type UserRepository interface {
	GetByID(ctx context.Context, id int) (*User, error)
	Save(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}

// User represents a user entity
type User struct {
	ID    int
	Name  string
	Email string
}

// UserService contains business logic for users
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, name, email string) (*User, error) {
	user := &User{
		Name:  name,
		Email: email,
	}
	if err := s.repo.Save(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser removes a user
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	// First check if user exists
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

