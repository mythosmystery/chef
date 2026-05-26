package projctx

// CompactResult summarizes compaction of a context file.
type CompactResult struct {
	File     FileName
	Before   int
	After    int
	Removed  int
}

// Compact triggers summarization to fit within budget.
func (m *Manager) Compact(name FileName) (CompactResult, error) {
	panic("not implemented")
}
