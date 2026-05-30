package config

import (
	"encoding/json"
	"fmt"
	"time"
)

// Duration wraps time.Duration for JSON unmarshaling of strings like "5m".
type Duration struct {
	time.Duration
}

// UnmarshalJSON accepts duration strings ("5m", "300s") or numeric nanoseconds.
func (d *Duration) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var n int64
	if err := json.Unmarshal(data, &n); err == nil {
		d.Duration = time.Duration(n)
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("duration: want string or number, got %s", string(data))
	}
	if s == "" {
		return fmt.Errorf("duration: empty string")
	}
	parsed, err := time.ParseDuration(s)
	if err != nil {
		return fmt.Errorf("duration: %w", err)
	}
	d.Duration = parsed
	return nil
}

// MarshalJSON writes duration as a string for readability.
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Duration.String())
}
