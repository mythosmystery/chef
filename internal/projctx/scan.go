package projctx

// DriftKind classifies a detected change.
type DriftKind string

const (
	DriftAdded   DriftKind = "added"
	DriftRemoved DriftKind = "removed"
	DriftChanged DriftKind = "changed"
)

// DriftItem describes one drift finding.
type DriftItem struct {
	Kind DriftKind
	Path string
}

// ScanReport summarizes project tree drift vs context.
type ScanReport struct {
	Items []DriftItem
}

// Scan compares the project tree against project.md entries.
func (m *Manager) Scan() (ScanReport, error) {
	panic("not implemented")
}
