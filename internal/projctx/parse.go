package projctx

// Entry is a keyed line in a context file.
type Entry struct {
	Key     string
	Content string
	Source  string
}

// Parse reads markdown content into entries.
func Parse(content string) ([]Entry, error) {
	panic("not implemented")
}
