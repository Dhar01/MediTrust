package models

import "time"

type User struct {
	Id         string
	Name       string
	Email      string
	Password   string
	Created_at time.Time
}
