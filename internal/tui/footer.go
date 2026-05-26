package tui

// FooterModel renders budget bar, token/cost stats, cwd, and model.
type FooterModel struct {
	CWD           string
	BudgetUsed    int
	BudgetMax     int
	InputTokens   int
	OutputTokens  int
	Cost          float64
	Model         string
	Width         int
}

// NewFooter creates a footer model.
func NewFooter(cwd string) FooterModel {
	return FooterModel{CWD: cwd}
}

// View renders the footer line.
func (f FooterModel) View() string {
	return renderFooterModel(f)
}

func renderFooterModel(f FooterModel) string {
	panic("not implemented")
}

// BudgetPercent returns budget usage as 0-100.
func (f FooterModel) BudgetPercent() int {
	if f.BudgetMax <= 0 {
		return 0
	}
	return (f.BudgetUsed * 100) / f.BudgetMax
}
