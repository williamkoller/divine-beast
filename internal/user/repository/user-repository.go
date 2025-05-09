package user_repository

import (
	"fmt"
	"sync"
)

type User struct {
	Email string
	Age   int
}

type UserRepository interface {
	AddUser(user User) bool
	GetUser(email string) (*User, bool)
}

type InMemoryUserRepository struct {
	users map[string]User
	mu    sync.RWMutex
}

var (
	instance *InMemoryUserRepository
	once     sync.Once
)

// GetInstance returns a singleton instance of InMemoryUserRepository
func GetInstance() *InMemoryUserRepository {
	once.Do(func() {
		instance = &InMemoryUserRepository{
			users: make(map[string]User),
		}
	})
	return instance
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return GetInstance()
}

func (repo *InMemoryUserRepository) AddUser(user User) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.users[user.Email]; exists {
		return false
	}

	repo.users[user.Email] = user
	return true
}

func (repo *InMemoryUserRepository) GetUser(email string) (*User, bool) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	user, exists := repo.users[email]
	if !exists {
		return nil, false
	}
	return &user, true
}

// String implements fmt.Stringer for better logging
func (repo *InMemoryUserRepository) String() string {
	return fmt.Sprintf("InMemoryUserRepository{users: %+v}", repo.users)
}
