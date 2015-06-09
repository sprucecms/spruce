package spruce

import (
	"testing"
)

func init() {
}

func TestNewSqlResourceRepo(t *testing.T) {
	r := NewSqlResourceRepo("http://localhost:5984")

	if r.address != "http://localhost:5984" {
		t.Fatalf("Sql repo address didn't get set correctly.")
	}
}
