package handlers

import (
	"encoding/json"
	"go-mux/domains"
	"go-mux/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

var userRepo = repositories.UserRepository{}

// UserIndex is for get user list
func UserIndex(w http.ResponseWriter, r *http.Request) {
	users, err := userRepo.All()
	if err != nil {
		res, _ := json.Marshal(map[string]string{"message": "user not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(map[string][]domains.User{"users": users})
	w.Write(res)
}

// UserShow for
func UserShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := userRepo.Find(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(map[string]domains.User{"user": user})
	w.Write(res)
}

// UserCreate is
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user domains.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdUser, _ := userRepo.Create(user)
	res, _ := json.Marshal(map[string]domains.User{"user": createdUser})
	w.Write(res)
}
