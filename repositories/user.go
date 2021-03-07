package repositories

import (
	"go-mux/domains"
)

// IUserRepository はユーザーの取得、保存、更新に関わるロジックをまとめたもの
type IUserRepository interface {
	Find(id string) (domains.User, error)
}

// UserRepository is
type UserRepository struct{}

var baseRepo = &BaseRepository{name: "users"}

var dummyUsers = []domains.User{
	{ID: "1", Name: "user1", Email: "user1@example.com", Birthday: "2000-01-01"},
	{ID: "2", Name: "user2", Email: "user2@example.com", Birthday: "2000-02-02"},
	{ID: "3", Name: "user3", Email: "user3@example.com", Birthday: "2000-03-03"},
}

// Find ユーザー単体を返す
func (UserRepository) Find(id string) (domains.User, error) {
	d, err := baseRepo.Find(id)
	return d.(domains.User), err
}

// All for
func (UserRepository) All() ([]domains.User, error) {
	d, err := baseRepo.All()
	return d.([]domains.User), err
}

// Create is
func (UserRepository) Create(user domains.User) (domains.User, error) {
	d, err := baseRepo.All()
	return d.(domains.User), err
}
