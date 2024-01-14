package models

import (
	"net/mail"
	"time"
)

type Email string
type Password string

func (p Password) IsValid() bool {
	//validation check
	if len(p) < 8 {
		return false
	}

	if len(p) > 64 {
		return false
	}

	return true
}

func (e Email) IsValid() bool {
	//validation check
	_, err := mail.ParseAddress(string(e))
	if err != nil {
		return false
	}

	return true
}

type User struct {
	UserID    string    `json:"user_id" bson:"user_id"`
	Username  string    `json:"username" bson:"username"`
	FullName  string    `json:"fullname" bson:"fullname"`
	Email     Email     `json:"email" bson:"email"`
	Password  Password  `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
