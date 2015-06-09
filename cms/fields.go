package cms

import null "gopkg.in/guregu/null.v2"

type Field interface {
	FieldName() string
	FieldLabel() string
	IsRequired() bool
}

type NamedField struct {
	Name        string
	Label       string
	Hint        string
	Placeholder string
	Description string
	Required    bool
}

func (n NamedField) FieldName() string {
	return n.Name
}

func (n NamedField) FieldLabel() string {
	return n.Label
}

func (n NamedField) FieldPlaceholder() string {
	return n.Placeholder
}

func (n NamedField) IsRequired() bool {
	return n.Required
}

type Text struct {
	NamedField
	DefaultValue *string
	MinLength    null.Int
	MaxLength    null.Int
}

func (t Text) IsValid() bool {
	return true
}

type Symbol struct {
	NamedField
	DefaultValue null.String
	ValidValues  []string
}

type Integer struct {
	NamedField
	DefaultValue null.Int
	MinValue     null.Int
	MaxValue     null.Int
}

type Decimal struct {
	NamedField
	DefaultValue null.Float
	MinValue     null.Float
	MaxValue     null.Float
}

type Boolean struct {
	NamedField
	DefaultValue bool
}

type Date struct {
	NamedField
	DefaultValue null.Bool
}

type Location struct {
	NamedField
	X      null.Float
	Y      null.Float
	Radius null.Float
}

type Entry struct {
	NamedField
	AllowedTypes []string
}

type ImageField struct { /* Rename to Image and put in a package to avoid collision with Image object */
	NamedField
	MaxBytes     int `json:",omitempty"`
	AllowedTypes []string
}

type Video struct {
	NamedField
	MaxBytes     int `json:",omitempty"`
	AllowedTypes []string
}
