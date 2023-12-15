package models

import "time"

type User struct {
	IDField
	Name     string `firestore:"name,omitempty" json:"name"`
	Email    string `firestore:"email,omitempty" json:"email"`
	Password string `firestore:"password,omitempty" json:"-"`
}

func (u *User) GetType() string {
	return "users"
}

func (u *User) GetIndex() string {
	return u.Email
}

type Token struct {
	IDField
	Key  string    `firestore:"key" json:"token"`
	User string    `firestore:"user" json:"user"`
	TTL  time.Time `firestore:"ttl" json:"expires"`
}

func (t *Token) GetType() string {
	return "tokens"
}
