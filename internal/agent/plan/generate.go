package plan

import "context"

// Generator creates plans from user prompts.
type Generator struct{}

// NewGenerator creates a plan generator.
func NewGenerator() *Generator {
	return &Generator{}
}

// Generate builds a structured plan from a prompt and context.
func (g *Generator) Generate(ctx context.Context, prompt string) (*Plan, error) {
	panic("not implemented")
}
