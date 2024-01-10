package models

type Like struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	PostId    string `json:"post_id"`
	CreatedAt string `json:"created_at"`
}
type LikeList struct {
	Likes []*Like `json:"like"`
	Count int     `json:"count"`
}
