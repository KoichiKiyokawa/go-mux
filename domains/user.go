package domains

// User は投稿者を表す
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
}
