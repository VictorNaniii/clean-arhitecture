package model

import "time"

type User struct {
	UUID      string
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserInfo struct {
	First_name string
	Last_name  string
	Age        int64
}
