package handlers

import (
	"encoding/json"
	"go-mux/domains"
	"go-mux/repositories"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// GetUsers is for get user list
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repositories.AllUser()
	if err != nil {
		res, _ := json.Marshal(map[string]string{"message": "user not found"})
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	res, _ := json.Marshal(map[string]domains.User{"user": user})
	w.Write(res)
}

// CreateUser is
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domains.User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)

	createdUser, _ := repositories.CreateUser(user)
	res, _ := json.Marshal(map[string]domains.User{"user": createdUser})
	w.Write(res)
}
