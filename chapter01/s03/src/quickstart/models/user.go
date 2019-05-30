package models

import (
	"time"
)

type User struct {
	ID       int
	UserName string
	Password string
	Birthday time.Time
	Addr     string
	Status   int
}
