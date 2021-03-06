package repositories

import (
	"errors"
	"go-mux/domains"
)

// UserRepository はユーザーの取得、保存、更新に関わるロジックをまとめたもの
type UserRepository interface {
	Find(id string) (domains.User, error)
}

var dummyUsers = []domains.User{
	{ID: "1", Name: "user1", Email: "user1@example.com", Birthday: "2000-01-01"},
	{ID: "2", Name: "user2", Email: "user2@example.com", Birthday: "2000-02-02"},
	{ID: "3", Name: "user3", Email: "user3@example.com", Birthday: "2000-03-03"},
}

// FindUser ユーザー単体を返す
func FindUser(id string) (domains.User, error) {
	for _, user := range dummyUsers {
		if user.ID == id {
			return user, nil
		}
	}

	var empty interface{}
	return empty.(domains.User), errors.New("user not found")
}

// AllUser for
func AllUser() ([]domains.User, error) {
	return dummyUsers, nil
}
