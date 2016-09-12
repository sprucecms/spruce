package sprucelib

type ContentTemplate struct {
	Name      string
	Content   []byte
	MediaType string // Default text/html, could also be application/xml or application/json
	IsRoot    bool   // For HTML templates, a "root" template is one that starts with <html>
}
