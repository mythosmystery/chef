package session

// Tree navigates parent/child session relationships.
type Tree struct {
	store *Store
}

// NewTree creates a session tree helper.
func NewTree(store *Store) *Tree {
	return &Tree{store: store}
}

// Fork creates a child session inheriting parent messages by reference.
func (t *Tree) Fork(parent Session) (Session, error) {
	panic("not implemented")
}

// Ancestors returns sessions from root to parent of id.
func (t *Tree) Ancestors(id string) ([]Session, error) {
	panic("not implemented")
}

// Messages returns the full message chain for a session id.
func (t *Tree) Messages(id string) ([]Message, error) {
	panic("not implemented")
}
