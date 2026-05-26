package workspace

import "fmt"

// SafetyError is returned when a path violates write restrictions.
type SafetyError struct {
	Path   string
	Reason string
}

func (e SafetyError) Error() string {
	return fmt.Sprintf("path blocked: %s (%s)", e.Path, e.Reason)
}

// CanWrite reports whether path is allowed for write/edit operations.
func (w *Workspace) CanWrite(path string) error {
	panic("not implemented")
}

// CanRead reports whether path is allowed for read operations.
func (w *Workspace) CanRead(path string) error {
	panic("not implemented")
}
