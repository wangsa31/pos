package utils

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateCredentials(email string, pass string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	claims := jwt.MapClaims{
		"email": email,
		"pass":  pass,
		// Token akan kadaluwarsa dalam 1 jam
	}

	// Create token jwt with example HMAC SHA-256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SCREET"))
	signedToken, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return signedToken, nil

}

func CheckCredentials(token_string string) (bool, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SCREET")), nil
	})
	if err != nil || !token.Valid {
		return false, err
	}
	return true, nil
}
