package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIssueFromJson(t *testing.T) {
	jsonString := `{
		"id":"1234", 
		"body":"body", 
		"created_at":"2022-05-02T15:04:05+02:00"
	}`

	issue, err := NewIssueFromJson([]byte(jsonString))

	assert.Nil(t, err)
	assert.NotNil(t, issue)
	assert.Equal(t, "1234", issue.ID)
	assert.Equal(t, "body", issue.Body)
	assert.Equal(t, "2022-05-02T15:04:05+02:00", issue.CreatedAt.Format("2006-01-02T15:04:05Z07:00"))

	jsonString = `{
		"id":"1234", 
		"body":"body", 
		"created_at":"2022T15:04:0+02:00"
	}`

	issue, err = NewIssueFromJson([]byte(jsonString))

	assert.NotNil(t, err)
	assert.Nil(t, issue)

}

func TestIssueToJson(t *testing.T) {
	jsonString := `{
		"id":"1234", 
		"body":"body", 
		"created_at":"2022-05-02T15:04:05+02:00"
	}`

	issue, err := NewIssueFromJson([]byte(jsonString))

	assert.Nil(t, err)
	assert.NotNil(t, issue)

	jsonBytes, err := IssueToJson(issue)

	assert.Nil(t, err)
	assert.NotNil(t, jsonBytes)

	jsonStringAfter := string(jsonBytes)

	assert.Equal(t, "{\"id\":\"1234\",\"body\":\"body\",\"created_at\":\"2022-05-02T15:04:05+02:00\"}", jsonStringAfter)
	// testlog.Println(jsonStringAfter)
}
