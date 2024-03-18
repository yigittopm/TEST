package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Sign(userId uint, expirationTime time.Duration) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(expirationTime).Unix()

	// Sign the token with a secret key
	secretKey := "your-secret-key"
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Verify(token string) (string, error) {
	// Verify the token and extract the claims
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Provide the secret key used for signing the token
		secretKey := "your-secret-key"
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	userId := claims["userId"].(string)
	return userId, nil
}
