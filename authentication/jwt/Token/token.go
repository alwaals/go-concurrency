package token

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return tokenString, nil
}

func VerifyToken(token string) error{
	t,err:=jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {
		return secretKey,nil
	})
	if err!=nil{
		return err
	}
	if !t.Valid{
		return errors.New("invalid token")
	}
	return nil
}