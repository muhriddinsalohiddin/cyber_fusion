package helper

import (
	"app/config"
	"app/models"

	"github.com/golang-jwt/jwt"
)

func NewAccessToken(claims models.UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(config.TOKEN_SECRET))
}

func ParseAccessToken(accessToken string) (*models.UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TOKEN_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedAccessToken.Claims.(*models.UserClaims), nil
}
