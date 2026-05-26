package provider

// Thinking configures extended reasoning for supported models.
type Thinking string

const (
	ThinkingOff    Thinking = "off"
	ThinkingLow    Thinking = "low"
	ThinkingMedium Thinking = "medium"
	ThinkingHigh   Thinking = "high"
)

// ParseThinking parses a thinking level string.
func ParseThinking(s string) Thinking {
	switch Thinking(s) {
	case ThinkingOff, ThinkingLow, ThinkingMedium, ThinkingHigh:
		return Thinking(s)
	default:
		return ThinkingMedium
	}
}
