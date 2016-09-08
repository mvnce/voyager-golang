package models

import "time"

type Comment struct {
	Id int64
	UserId int64
	PostId int64
	Content string
	Status string
	Created time.Time
	Updated time.Time
}
