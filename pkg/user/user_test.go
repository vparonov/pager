package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCtor(t *testing.T) {
	jsonString := `{"id":"admin", "name":"Admin Admin", "email":"admin@example.com"}`

	user, err := NewFromJSON([]byte(jsonString))

	assert.Nil(t, err)

	assert.Equal(t, "admin", user.ID)
	assert.Equal(t, "Admin Admin", user.Name)
	assert.Equal(t, "admin@example.com", user.EMail)

	marshaledUser, err := user.ToJSON()

	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(marshaledUser))

	unMarshaledUser, err := NewFromJSON(marshaledUser)

	assert.Nil(t, err)

	assert.Equal(t, user.ID, unMarshaledUser.ID)
	assert.Equal(t, user.Name, unMarshaledUser.Name)
	assert.Equal(t, user.EMail, unMarshaledUser.EMail)
}
