// Package versions provides types for the Headscale API.
package versions

import "fmt"

// APIVersion represents the API version.
type APIVersion string

const (
	// APIVersionV1 is the v1 API version.
	APIVersionV1 APIVersion = "v1"
)

// String returns the string representation of the API version.
func (v APIVersion) String() string {
	return string(v)
}

// Validate validates the API version.
func (v APIVersion) Validate() error {
	switch v {
	case APIVersionV1:
		return nil
	default:
		return fmt.Errorf("invalid API version: %s", v)
	}
}

// GetBasePath returns the base path for the API version.
func (v APIVersion) GetBasePath() string {
	return fmt.Sprintf("/api/%s", v.String())
}
