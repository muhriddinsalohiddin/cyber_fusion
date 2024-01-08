package models

type Author struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type AuthorList struct {
	Authors []*Author `json:"authors"`
	Count   int       `json:"count"`
}
