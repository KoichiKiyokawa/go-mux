package handlers

import (
	"encoding/json"
	"go-mux/domains"
	"go-mux/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

// GetUsers is for get user list
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repositories.AllUser()
	if err != nil {
		res, _ := json.Marshal(map[string]string{"message": "user not found"})
		w.Write(res)
		return
	}

	res, _ := json.Marshal(map[string][]domains.User{"users": users})
	w.Write(res)
}

// GetUser for
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := repositories.FindUser(params["id"])
	if err != nil {
		res, _ := json.Marshal(map[string]string{"message": "user not found"})
		w.Write(res)
		return
	}
	res, _ := json.Marshal(map[string]domains.User{"user": user})
	w.Write(res)
}
