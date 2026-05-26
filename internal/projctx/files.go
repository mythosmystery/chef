package projctx

// FileName identifies a known context file.
type FileName string

const (
	FileProject      FileName = "project.md"
	FileGlossary     FileName = "glossary.md"
	FileFeatures     FileName = "features.md"
	FileConventions  FileName = "conventions.md"
	FileArchitecture FileName = "architecture.md"
	FileData         FileName = "data.md"
	FileAPI          FileName = "api.md"
	FileWorkflows    FileName = "workflows.md"
)

// KnownFiles returns all supported context file names.
func KnownFiles() []FileName {
	return []FileName{
		FileProject, FileGlossary, FileFeatures, FileConventions,
		FileArchitecture, FileData, FileAPI, FileWorkflows,
	}
}

// BudgetFor returns the token budget for a file name.
func (m *Manager) BudgetFor(name FileName) int {
	if m.cfg.Budget == nil {
		return 0
	}
	return m.cfg.Budget[string(name)]
}
