package pkg

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	privKeyPath = "pkg/keys/private.pem"
	pubKeyPath  = "pkg/keys/public.pem"
)

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
	serverPort int
)

func InitializeKeys() error {
	// Load private key
	privateKeyPEM, err := os.ReadFile(privKeyPath)
	if err != nil {
		return err
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return err
	}

	// Load public key
	publicKeyPEM, err := os.ReadFile(pubKeyPath)
	if err != nil {
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		return err
	}

	return nil
}

type Claims struct {
	jwt.RegisteredClaims
	User UserInfo
}

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Token struct {
	Token     string `json:"access_token"`
	ExpiresAt string `json:"expires_at"`
}

func GenerateToken(userID, email string) (Token, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "golang-book.beneboba.me",
			Subject:   userID,
		},
		User: UserInfo{
			ID:    userID,
			Email: email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	accessToken, err := token.SignedString(signKey)

	if err != nil {
		return Token{}, err
	}

	return Token{
		Token:     accessToken,
		ExpiresAt: expirationTime.Format(time.DateTime),
	}, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
