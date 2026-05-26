package refs

// Attachment is resolved content ready to inject into a prompt.
type Attachment struct {
	Reference Reference
	Content   string
	Tokens    int
	Truncated bool
}

// Attach resolves references and loads their content.
func (r *Resolver) Attach(refs []Reference) ([]Attachment, error) {
	panic("not implemented")
}

// AttachOne resolves a single reference.
func (r *Resolver) AttachOne(ref Reference) (Attachment, error) {
	panic("not implemented")
}
