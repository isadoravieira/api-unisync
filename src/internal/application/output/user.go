package output

import "time"

type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	UserName  string    `json:"userName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type UserStore struct {
	Data User `json:"data"`
}

type UserOutput struct {
}
