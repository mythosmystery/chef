// Package diff provides git and last-edit diff utilities.
package diff

// Result holds diff output for one or more files.
type Result struct {
	Path string
	Text string
}

// Service computes diffs for the diff tool.
type Service struct{}

// New creates a diff service.
func New() *Service {
	return &Service{}
}

// GitDiff returns uncommitted git diff for path (empty path = all).
func (s *Service) GitDiff(path string) ([]Result, error) {
	panic("not implemented")
}

// AllGitDiff returns all uncommitted changes.
func (s *Service) AllGitDiff() ([]Result, error) {
	panic("not implemented")
}

// LastEditDiff returns the last edit diff tracked by the tool layer.
func (s *Service) LastEditDiff(path string) (Result, error) {
	panic("not implemented")
}
