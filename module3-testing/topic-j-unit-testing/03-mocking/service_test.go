// Assignment 3: Mocking with Interfaces
//
// Goal: Create a mock implementation of UserRepository for testing.
//
// Run: go test -v ./...

package service

import (
	"context"
	"errors"
	"testing"
)

// MockUserRepository is a test double for UserRepository
type MockUserRepository struct {
	users      map[int]*User
	saveError  error
	getError   error
	saveCalled int
	getCalled  int
}

// NewMockUserRepository creates a new mock repository
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[int]*User),
	}
}

func (m *MockUserRepository) GetByID(ctx context.Context, id int) (*User, error) {
	m.getCalled++
	if m.getError != nil {
		return nil, m.getError
	}
	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) Save(ctx context.Context, user *User) error {
	m.saveCalled++
	if m.saveError != nil {
		return m.saveError
	}
	// Auto-increment ID
	if user.ID == 0 {
		user.ID = len(m.users) + 1
	}
	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) Delete(ctx context.Context, id int) error {
	delete(m.users, id)
	return nil
}

// Helper methods for test setup
func (m *MockUserRepository) SetGetError(err error)  { m.getError = err }
func (m *MockUserRepository) SetSaveError(err error) { m.saveError = err }
func (m *MockUserRepository) AddUser(user *User)     { m.users[user.ID] = user }

// Tests

func TestUserService_GetUser(t *testing.T) {
	// Setup
	mockRepo := NewMockUserRepository()
	mockRepo.AddUser(&User{ID: 1, Name: "Alice", Email: "alice@test.com"})

	service := NewUserService(mockRepo)
	ctx := context.Background()

	// Test successful get
	t.Run("existing user", func(t *testing.T) {
		user, err := service.GetUser(ctx, 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if user.Name != "Alice" {
			t.Errorf("expected name Alice, got %s", user.Name)
		}
	})

	// Test not found
	t.Run("non-existing user", func(t *testing.T) {
		_, err := service.GetUser(ctx, 999)
		if err == nil {
			t.Error("expected error for non-existing user")
		}
	})
}

func TestUserService_CreateUser(t *testing.T) {
	t.Run("successful creation", func(t *testing.T) {
		mockRepo := NewMockUserRepository()
		service := NewUserService(mockRepo)
		ctx := context.Background()

		user, err := service.CreateUser(ctx, "Bob", "bob@test.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if user.Name != "Bob" {
			t.Errorf("expected name Bob, got %s", user.Name)
		}
		if mockRepo.saveCalled != 1 {
			t.Errorf("expected Save to be called once, called %d times", mockRepo.saveCalled)
		}
	})

	t.Run("save error", func(t *testing.T) {
		mockRepo := NewMockUserRepository()
		mockRepo.SetSaveError(errors.New("database error"))
		service := NewUserService(mockRepo)
		ctx := context.Background()

		_, err := service.CreateUser(ctx, "Bob", "bob@test.com")
		if err == nil {
			t.Error("expected error when save fails")
		}
	})
}

func TestUserService_DeleteUser(t *testing.T) {
	t.Run("delete existing user", func(t *testing.T) {
		mockRepo := NewMockUserRepository()
		mockRepo.AddUser(&User{ID: 1, Name: "Alice"})
		service := NewUserService(mockRepo)
		ctx := context.Background()

		err := service.DeleteUser(ctx, 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("delete non-existing user", func(t *testing.T) {
		mockRepo := NewMockUserRepository()
		service := NewUserService(mockRepo)
		ctx := context.Background()

		err := service.DeleteUser(ctx, 999)
		if err == nil {
			t.Error("expected error when deleting non-existing user")
		}
	})
}

