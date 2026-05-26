package session

// ResumeOptions selects how to resume a session.
type ResumeOptions struct {
	Continue bool
	Browse   bool
	Path     string
}

// Resume resolves the session to load based on CLI options.
func (s *Store) Resume(workingDir string, opts ResumeOptions) (Session, []Message, error) {
	panic("not implemented")
}

// MostRecent returns the latest session for workingDir.
func (s *Store) MostRecent(workingDir string) (Session, error) {
	panic("not implemented")
}
