package diff

// Tracker records last edit operations for diff --last.
type Tracker struct{}

// NewTracker creates an edit tracker.
func NewTracker() *Tracker {
	return &Tracker{}
}

// Record stores the before/after state of an edit.
func (t *Tracker) Record(path, before, after string) {
	panic("not implemented")
}

// Last returns the last recorded diff for path.
func (t *Tracker) Last(path string) (string, error) {
	panic("not implemented")
}
