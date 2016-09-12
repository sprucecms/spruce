package sprucelib

// A Field is a concrete instance of a FieldType attached to a ContentType in a particular
// position.
//
// For example, the "Department" Field on a "Staff" ContentType which should be the
// 3rd field in the UI.
//
// NOTE: This type does not store data values. In the example above, the value "Marketing"
// is attached to a Resource as a FieldValue
type Field struct {
	ID        string // TODO Should this be a GUID?
	Name      string
	TypeID    string
	Type      FieldType
	Position  int // Used to sort fields in the UI
	MinValues int // The minimum number of FieldValues required to be assigned in the UI
	MaxValues int // The maximum number of FieldValues allowed to be assigned in the UI
}
