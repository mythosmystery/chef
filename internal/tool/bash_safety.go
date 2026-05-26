package tool

// SafetyCheck validates a bash command against the blocklist.
type SafetyCheck struct {
	Blocklist []string
}

// RequiresConfirmation reports whether command needs user approval.
func (s SafetyCheck) RequiresConfirmation(command string) (bool, string) {
	panic("not implemented")
}

// IsBlocked reports whether command is unconditionally blocked.
func (s SafetyCheck) IsBlocked(command string) (bool, string) {
	panic("not implemented")
}
