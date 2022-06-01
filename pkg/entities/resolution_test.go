package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResolutionFromJSON(t *testing.T) {
	jsonString := `{
		"issue_id":"1234", 
		"user_id":"vangel@elastecad.com", 
		"ts":"2022-05-02T15:04:05+02:00"
	}`

	resolution, err := NewResolutionFromJson([]byte(jsonString))

	assert.Nil(t, err)
	assert.NotNil(t, resolution)

	jsonBytes, err := ResolutionToJson(resolution)

	assert.Nil(t, err)
	assert.NotNil(t, jsonBytes)

	jsonStringAfter := string(jsonBytes)

	assert.Equal(t, "{\"issue_id\":\"1234\",\"user_id\":\"vangel@elastecad.com\",\"ts\":\"2022-05-02T15:04:05+02:00\"}", jsonStringAfter)

}
