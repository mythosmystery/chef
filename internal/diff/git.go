package diff

// Git runs git diff commands for uncommitted changes.
type Git struct{}

// DiffPath returns git diff for a single path.
func (Git) DiffPath(path string) (string, error) {
	panic("not implemented")
}

// DiffAll returns git diff for all uncommitted changes.
func (Git) DiffAll() (string, error) {
	panic("not implemented")
}
