package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("TRITONSTORAGE_SECRETKEY")

type Credentials struct {
	Password string 
	Username string
	Email string
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

func CreateToken(username string, password string, email string) (time.Time, string, int) {
	
	expirationTime := time.Now().Add(30*time.Minute)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err!= nil {
		// Return 1 something error happen
		return expirationTime,"",1
	}

	// Return 0 successful
	return expirationTime, tokenString, 0
}