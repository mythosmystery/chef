// Package tokens provides token counting utilities.
package tokens

// Counter estimates token usage for text.
type Counter interface {
	Count(text string) int
}

// EstimateCounter uses character count / 4 as a token estimate.
type EstimateCounter struct{}

// Count returns an estimated token count for text.
func (EstimateCounter) Count(text string) int {
	return Estimate(text)
}
