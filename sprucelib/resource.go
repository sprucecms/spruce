package sprucelib

import (
	"time"
)

type Resource struct {
	Id          uint
	Alias       string
	URI         string
	TreePath    string
	Title       string
	ShortTitle  string
	IsPublished bool
	PublishAt   time.Time
	UnpublishAt time.Time
	PublishedBy *User
}
