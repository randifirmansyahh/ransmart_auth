package tokenHelper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	AUD            = os.Getenv("JWT_AUD")
	ISS            = os.Getenv("JWT_ISS")
	LOGIN_SECRET   = os.Getenv("JWT_SECRET_KEY")
	WAKTU          = os.Getenv("JWT_EXPIRATION_DURATION_DAY")
	MESSAGE        = "Unathorized!"
	MESSAGE_KOSONG = "Requst authorize kosong"
)

func BuatJWT(iss, aud, secret string, exp time.Time) (string, error) {
	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    iss,
		"Audience":  aud,
		"ExpiresAt": exp.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	// membuat jwtnya
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
