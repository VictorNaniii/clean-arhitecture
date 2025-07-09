package model

import "time"

type user struct {
	UUID      string
	info      UserInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserInfo struct {
	first_name string
	last_name  string
	age        int64
}
