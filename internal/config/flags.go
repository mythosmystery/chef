package config

import (
	"strings"
)

// FlagOverrides holds CLI flag overrides applied after file config.
type FlagOverrides struct {
	Provider  string
	Model     string
	Thinking  string
	Tools     string
	NoTools   bool
	NoContext bool
}

// ApplyFlags overlays CLI flags onto cfg.
func ApplyFlags(cfg *Config, flags FlagOverrides) error {
	if flags.Provider != "" {
		cfg.Provider = flags.Provider
	}
	if flags.Model != "" {
		cfg.Model = flags.Model
	}
	if flags.Thinking != "" {
		if err := ValidateThinkingFlag(flags.Thinking); err != nil {
			return err
		}
		cfg.Thinking = flags.Thinking
	}
	if flags.NoTools {
		cfg.Tools = nil
	} else if flags.Tools != "" {
		cfg.Tools = splitTools(flags.Tools)
	}
	return nil
}

func splitTools(list string) []string {
	parts := strings.Split(list, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// NoContextFromFlags reports whether context injection should be skipped.
func NoContextFromFlags(flags FlagOverrides) bool {
	return flags.NoContext
}
