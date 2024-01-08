package models

type Books struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type BooksList struct {
	Books []*Books `json:"books"`
	Count int      `json:"count"`
}

type LsitBookReq struct {
	Title string
	Author string
	Description string
}

