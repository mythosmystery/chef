package config

import "fmt"

// ErrGlobalConfigMissing is returned when RequireGlobal is set and the global
// config file does not exist.
type ErrGlobalConfigMissing struct {
	Path string
}

func (e ErrGlobalConfigMissing) Error() string {
	return fmt.Sprintf(`chef is not configured.

No global config file found at %s.

Run the setup wizard to create one:
  chef config

For project-specific settings:
  chef config --project
`, e.Path)
}
