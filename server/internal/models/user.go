package models

import "time"

type User struct {
	Id         int
	Username   string
	Password   string
	Created_At *time.Time
	Updated_At *time.Time
}

func (User) TableName() string {
	return "users"
}
