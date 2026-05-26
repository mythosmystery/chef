package projctx

import "sync"

// Store handles disk I/O for context files with per-file locking.
type Store struct {
	dir   string
	locks map[FileName]*sync.Mutex
}

// NewStore creates a context file store.
func NewStore(dir string) *Store {
	return &Store{
		dir:   dir,
		locks: make(map[FileName]*sync.Mutex),
	}
}

// Read loads a context file from disk.
func (s *Store) Read(name FileName) (string, error) {
	panic("not implemented")
}

// Write saves a context file to disk.
func (s *Store) Write(name FileName, content string) error {
	panic("not implemented")
}

// Exists reports whether the context directory is initialized.
func (s *Store) Exists() (bool, error) {
	panic("not implemented")
}
