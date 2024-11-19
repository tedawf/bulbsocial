package store

import (
	"context"
	"database/sql"
	"time"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct{}

func (m *MockUserStore) Create(context.Context, *sql.Tx, *User) error {
	return nil
}
func (m *MockUserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	return &User{ID: userID}, nil
}
func (m *MockUserStore) CreateAndInvite(context.Context, *User, string, time.Duration) error {
	return nil
}
func (m *MockUserStore) Verify(context.Context, string) error {
	return nil
}
func (m *MockUserStore) Delete(context.Context, int64) error {
	return nil
}
func (m *MockUserStore) GetByEmail(context.Context, string) (*User, error) {
	return &User{}, nil
}
