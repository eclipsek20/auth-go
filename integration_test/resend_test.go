package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eclipsek20/auth-go/types"
)

func TestResend(t *testing.T) {
	assert := assert.New(t)

	email := randomEmail()
	err := client.Resend(types.ResendRequest{
		Type:  types.ResendTypeSignup,
		Email: email,
	})
	assert.NoError(err)

	err = client.Resend(types.ResendRequest{})
	assert.Error(err)
}
