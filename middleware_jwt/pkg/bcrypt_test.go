package pkg_test

import (
	"middleware_jwt/pkg"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"

	hashedPassword, err := pkg.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	require.NotEqual(t, password, hashedPassword)

}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	tests := []struct {
		name     string
		password string
		hash     string
		expected bool
	}{
		{
			name:     "Correct password",
			password: password,
			hash:     string(hashedPassword),
			expected: true,
		},
		{
			name:     "Incorrect password",
			password: "wrongpassword",
			hash:     string(hashedPassword),
			expected: false,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     string(hashedPassword),
			expected: false,
		},
		{
			name:     "Empty hash",
			password: password,
			hash:     "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pkg.CheckPasswordHash(tt.password, tt.hash)
			require.Equal(t, tt.expected, result)
		})
	}
}
