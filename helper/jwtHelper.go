package helper

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func (c *JWTClaims) IsAdmin() bool {
	return c.Role == "admin"
}

func init() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func GetJWTSecret() string {
	return string(jwtSecret)
}

func GenerateJWT(user_id, role string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Second)

	claims := &JWTClaims{
		UserID: user_id,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetClaimsFromToken(tokenString string) (*JWTClaims, error) {
	token, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func extractToken(tokenString string) string {
	if len(tokenString) > 6 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	return tokenString
}
