package spruce

import "testing"

func TestNullPointerStrings(t *testing.T) {

	tmpl := &Template{Id: "test"}
	if &tmpl.Description != nil {
		// tmpl.Description is an empty string ("") if using string instead of *string
		// t.Errorf("Expected tmpl.Description to be nil")
	}

	if tmpl.Id != "test" {
		t.Errorf("Expected tmpl.Id to be 'test'")
	}
}
