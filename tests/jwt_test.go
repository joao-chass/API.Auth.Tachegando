package tests

import (
	"AuthTachegando/pkg/jwt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")

	token, err := jwt.GenerateJWT("usuarioTeste")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}
