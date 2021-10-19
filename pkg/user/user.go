package user

import "encoding/json"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	EMail string `json:"email"`
}

func NewFromJSON(data []byte) (*User, error) {
	var user User
	err := json.Unmarshal(data, &user)

	return &user, err
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
