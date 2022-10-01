package entities

import (
	"encoding/json"
	"time"
)

type Resolution struct {
	IssueID string    `json:"issue_id"`
	UserID  string    `json:"user_id"`
	Ts      time.Time `json:"ts"`
}

func NewResolutionFromJson(b []byte) (*Resolution, error) {
	var newResolution Resolution
	err := json.Unmarshal(b, &newResolution)

	if err != nil {
		return nil, err
	}
	return &newResolution, nil
}

func NewResolutionsSliceFromJson(b []byte) ([]*Resolution, error) {
	var newResolutionsSlice []*Resolution
	err := json.Unmarshal(b, &newResolutionsSlice)

	if err != nil {
		return nil, err
	}
	return newResolutionsSlice, nil
}

func ResolutionToJson(resolution *Resolution) ([]byte, error) {
	return json.Marshal(resolution)
}

func ResolutionSliceToJson(resolutions []*Resolution) ([]byte, error) {
	return json.Marshal(resolutions)
}
