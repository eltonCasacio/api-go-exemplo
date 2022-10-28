package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("elton", "email@hotmail.com", "123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Name)
	assert.Equal(t, "elton", user.Name)
	assert.Equal(t, "email@hotmail.com", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("elton", "email@hotmail.com", "123")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123"))
	assert.False(t, user.ValidatePassword("1237"))
	assert.NotEqual(t, "123", user.Password)
}
