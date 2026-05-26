package session

import "io"

// Store reads and writes JSONL session files.
type Store struct {
	dir string
}

// NewStore creates a session store rooted at dir.
func NewStore(dir string) *Store {
	return &Store{dir: dir}
}

// Save persists messages for a session.
func (s *Store) Save(sess Session, messages []Message) error {
	panic("not implemented")
}

// Load reads a session and its messages.
func (s *Store) Load(path string) (Session, []Message, error) {
	panic("not implemented")
}

// List returns sessions for a working directory.
func (s *Store) List(workingDir string) ([]Session, error) {
	panic("not implemented")
}

// OpenWriter returns a writer for streaming JSONL append.
func (s *Store) OpenWriter(sess Session) (io.WriteCloser, error) {
	panic("not implemented")
}
