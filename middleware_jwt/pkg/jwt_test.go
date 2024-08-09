package pkg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateTestKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func writeTestKeysToFile(t *testing.T, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) {
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	assert.NoError(t, err)

	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)

	err = os.WriteFile(privKeyPath, privateKeyPEM, 0600)
	assert.NoError(t, err)

	err = os.WriteFile(pubKeyPath, publicKeyPEM, 0600)
	assert.NoError(t, err)
}

func TestInitializeKeys(t *testing.T) {
	privateKey, publicKey, err := generateTestKeys()
	assert.NoError(t, err)

	writeTestKeysToFile(t, privateKey, publicKey)
	defer os.Remove(privKeyPath)
	defer os.Remove(pubKeyPath)

	err = InitializeKeys()
	assert.NoError(t, err)
	assert.NotNil(t, signKey)
	assert.NotNil(t, verifyKey)
}

func TestGenerateToken(t *testing.T) {
	privateKey, publicKey, err := generateTestKeys()
	assert.NoError(t, err)

	signKey = privateKey
	verifyKey = publicKey

	token, err := GenerateToken("12345", "user@example.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, token.Token)
	assert.NotEmpty(t, token.ExpiresAt)
}

func TestValidateToken(t *testing.T) {
	privateKey, publicKey, err := generateTestKeys()
	assert.NoError(t, err)

	signKey = privateKey
	verifyKey = publicKey

	token, err := GenerateToken("12345", "user@example.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, token.Token)

	claims, err := ValidateToken(token.Token)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, "12345", claims.Subject)
	assert.Equal(t, "user@example.com", claims.User.Email)
}

func TestValidateTokenInvalid(t *testing.T) {
	_, _, err := generateTestKeys()
	assert.NoError(t, err)

	_, err = ValidateToken("invalidToken")
	assert.Error(t, err)
}
