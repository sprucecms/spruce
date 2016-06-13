package sprucelib

import (
	"time"
)

type User struct {
	ID                    uint `gorm:"primary_key"`
	Username              string
	Email                 string
	PasswordHash          string
	PasswordResetToken    string
	FirstName             string
	LastName              string
	IsBlocked             bool
	Tokens                []UserToken
	CreatedAt             time.Time
	UpdatedAt             time.Time
	PasswordLastChangedAt time.Time
	LastLoginAt           time.Time
	FailedLoginAttempts   int
	DeletedAt             *time.Time
}
