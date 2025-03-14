package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTTools handles JWT operations
type JWTTools struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewJWTTools initializes JWT tools with key pairs
func NewJWTTools(privateKeyPath, publicKeyPath string) (*JWTTools, error) {
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, err
	}

	publicKeyData, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, err
	}

	return &JWTTools{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// GenerateToken creates a new JWT token
func (j *JWTTools) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(j.privateKey)
}

// ValidateToken verifies a JWT token
func (j *JWTTools) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return j.publicKey, nil
	})
}
