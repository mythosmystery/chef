package projctx

// BudgetUsage reports token usage for one file.
type BudgetUsage struct {
	File     FileName
	Used     int
	Max      int
	Overflow bool
}

// BudgetReport summarizes usage across all context files.
type BudgetReport struct {
	Files []BudgetUsage
	Total int
	Max   int
}

// Budget returns current token usage per file and total.
func (m *Manager) Budget() (BudgetReport, error) {
	panic("not implemented")
}

// ValidateUpdate checks whether an update would exceed budget.
func (m *Manager) ValidateUpdate(name FileName, content string) error {
	panic("not implemented")
}
