package models

import "github.com/golang-jwt/jwt"

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
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
	Token         string          `json:"token"`
	Comments      []*Comment      `json:"comments"`
	Notifications []*Notification `json:"notifications"`
	Posts         []*Post         `json:"posts"`
	Message       []*Message      `json:"messages"`
	Likes         []*Like         `json:"likes"`
	Claimes       *UserClaims     `json:"claimes"`
}

type UserClaims struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Duration int64  `json:"duration"`
	jwt.StandardClaims
}

type Users struct {
	Users []*User `json:"users"`
}

type UserReq struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
	Login    string `json:"login"`
}
