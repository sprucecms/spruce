package sprucelib

type ContentType struct {
	Id              int
	Name            string
	DataFields      []Field
	TitleAllowed    bool
	TitleRequired   bool
	ContentAllowed  bool
	ContentRequired bool
}
