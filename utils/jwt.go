package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func LoadEnv() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("error to load the key")
	}
	return os.Getenv("KEY"), err

}

func GenereteToken(email string, userId int64) (string, error) {
	secretKey, err := LoadEnv()
	if err != nil {
		return "", errors.New("error to load the key")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))

}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		secretKey, err := LoadEnv()
		if err != nil {
			return "", errors.New("error to load the key")
		}
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenISValid := parsedToken.Valid

	if !tokenISValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
