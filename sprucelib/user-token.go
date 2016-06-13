package sprucelib

import (
	"time"
)

type UserToken struct {
	UserID    uint   `gorm:"primary_key"`
	Token     string `gorm:"primary_key"`
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}
