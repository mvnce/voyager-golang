package models

import "time"

type Post struct {
	Id 	int64
	UserId 	int64
	Title 	string
	Content string
	Status 	string
	Created time.Time
	Updated time.Time
}