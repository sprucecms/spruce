package sprucelib

import ()

type FieldValue struct {
	Field Field
	Value interface{}
}

func (fv *FieldValue) IsValid() bool {
	return fv.Value != nil
}
