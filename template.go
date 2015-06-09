package spruce

import (
	"time"
)

type Template struct {
	Id              string
	Description     string
	IsRoot          bool
	CreatedAt       time.Time
	CreatedByUserId string
	UpdatedAt       time.Time
	UpdatedByUserId string
	CacheTTL        int
}

type TemplateService interface {

	// Retrieve a single template with a known id
	Get(id string) (*Template, error)

	// Submits a Template to the server. If a template with the given
	// id didn't already exist, it will be created, and the created
	// return value will be true.
	Store(template *Template) (created bool, err error)
}
