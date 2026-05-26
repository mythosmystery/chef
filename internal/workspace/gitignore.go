package workspace

// Gitignore provides .gitignore-aware path filtering.
type Gitignore struct {
	root string
}

// NewGitignore loads gitignore rules for root.
func NewGitignore(root string) (*Gitignore, error) {
	panic("not implemented")
}

// Match reports whether path should be ignored.
func (g *Gitignore) Match(path string) bool {
	panic("not implemented")
}
