package models

type Comment struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	UserId      string `json:"user_id"`
}
