package models

type User struct {
	Id            string          `json:"id"`
	Name          string          `json:"name"`
	Gender        bool            `json:"gender"`
	Birthday      string          `json:"birthday"`
	Email         string          `json:"email"`
	Login         string          `json:"login"`
	Password      string          `json:"password"`
	Bio           string          `json:"bio"`
	CreatedAt     string          `json:"created_at"`
	UpdetadAt     string          `json:"updetad_at"`
	Messages      []*Message      `json:"messages"`
	Notifications []*Notification `json:"notifications"`
	Posts         []*Post         `json:"posts"`
}

type Users struct {
	Users []*User `json:"users"`
}

type UserReq struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
