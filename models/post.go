package models

type Post struct {
	Id        string     `json:"id"`
	UserId    string     `json:"user_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	Comments  []*Comment `json:"comments"`
	Likes     []*Like    `json:"likes"`
}

type Like struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	PostId    string `json:"post_id"`
	CreatedAt string `json:"created_at"`
}
type PostListResp struct {
	Post  []*Post `json:"post"`
	Count int     `json:"count"`
}
type PostListReq struct {
	Post  []*Post `json:"post"`
	Count int     `json:"count"`
}
