package models


type Comment struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	PostId    string    `json:"post_id"`
	Body      string    `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type List struct {
	Comments []*Comment `json:"comment"`
	Count    int        `json:"count"`
}
