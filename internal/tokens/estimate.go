package tokens

// Estimate returns an estimated token count using chars/4.
func Estimate(text string) int {
	if len(text) == 0 {
		return 0
	}
	return (len(text) + 3) / 4
}
