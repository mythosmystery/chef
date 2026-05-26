package refs

// Completion is a suggestion for @ reference autocomplete.
type Completion struct {
	Label string
	Value string
	Kind  Kind
}

// Complete returns autocomplete suggestions for partial @ input.
func (r *Resolver) Complete(partial string) ([]Completion, error) {
	panic("not implemented")
}
