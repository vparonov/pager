package entities

import (
	"encoding/json"
	"time"
)

type Issue struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func NewIssueFromJson(b []byte) (*Issue, error) {
	var newIssue Issue
	err := json.Unmarshal(b, &newIssue)

	if err != nil {
		return nil, err
	}
	return &newIssue, nil
}

func IssueToJson(issue *Issue) ([]byte, error) {
	return json.Marshal(issue)
}
