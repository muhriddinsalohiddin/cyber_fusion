package models

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	Birthday  string `json:"birthday"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
