package models

type Like struct {
	id         string `json:"like_id"`
	user_id    string `json:"user_id"`
	post_id    string `json:"post_id"`
	created_at string `json"created_at"`
}
