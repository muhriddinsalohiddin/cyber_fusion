package models
<<<<<<< HEAD

type Comment struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	PostId    string `json:"post_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type List struct {
	Comments []*Comment `json:"comment"`
	Count    int        `json:"count"`
}
=======
>>>>>>> d9baf0fd15d6d098706e618c0340843b65bb3f2c
