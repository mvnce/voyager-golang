package models

import "time"

type User struct {
	Id 	int64
	Name 	string
	Password string
	Salt 	string
	Email 	string
	Type 	string
	Joined 	time.Time
}
