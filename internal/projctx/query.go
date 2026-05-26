package projctx

// QueryResult is a matched context snippet with attribution.
type QueryResult struct {
	File    FileName
	Key     string
	Snippet string
	Score   float64
}

// Query searches context files for topic matches.
func (m *Manager) Query(topic string) ([]QueryResult, error) {
	panic("not implemented")
}
