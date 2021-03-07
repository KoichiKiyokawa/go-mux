package handlers

import (
	"encoding/json"
	"go-mux/domains"
	"net/http"
	"time"

	"github.com/go-playground/validator"
)

// PostInput は登録, 更新処理のリクエストのバリデーションを定義する
type PostInput struct {
	Title string `validate:"required"`
	Body  string `validate:"required"`
}

func (p PostInput) getPostDomain() domains.Post {
	return domains.Post{
		Title:     p.Title,
		Body:      p.Body,
		CreatedAt: time.Now(),
	}
}

// PostCreate is
func PostCreate(w http.ResponseWriter, r *http.Request) {
	var postInput PostInput
	if err := json.NewDecoder(r.Body).Decode(&postInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(map[string]string{"msg": "bad request"})
		w.Write(res)
		return
	}

	validate := validator.New()
	if err := validate.Struct(postInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(map[string]string{"msg": "bad schema"})
		w.Write(res)
		return
	}

	var post domains.Post
	post = postInput.getPostDomain()
	res, _ := json.Marshal(post)
	w.Write(res)
}
