package command

import "strings"

// Match represents a slash command found anywhere in input.
type Match struct {
	Name string
	Args []string
	Start int
	End   int
}

// Parse finds slash commands anywhere in input text.
func Parse(input string) ([]Match, error) {
	panic("not implemented")
}

// ExtractCommand returns the first slash command match in input.
func ExtractCommand(input string) (Match, bool) {
	matches, err := Parse(input)
	if err != nil || len(matches) == 0 {
		return Match{}, false
	}
	return matches[0], true
}

// Strip removes slash command text from input, leaving free text.
func Strip(input string, match Match) string {
	before := strings.TrimSpace(input[:match.Start])
	after := strings.TrimSpace(input[match.End:])
	return strings.TrimSpace(before + " " + after)
}
