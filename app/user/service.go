package user

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken ...
func GenerateToken(name string, admin bool) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow", // Name
		true,       // Admin
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte("secret"))
}
