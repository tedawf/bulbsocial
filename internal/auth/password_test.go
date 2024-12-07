package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tedawf/bulbsocial/internal/util"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := util.RandomString(10)

	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckPassword(hashedPassword1, password)
	require.NoError(t, err)

	wrongPassword := util.RandomString(11)
	err = CheckPassword(hashedPassword1, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
