package versions

import (
	"fmt"
	"testing"
)

func TestAPIVersion_String(t *testing.T) {
	cases := []struct {
		name     string
		version  APIVersion
		expected string
	}{
		{"v1", APIVersionV1, "v1"},
		{"unknown", APIVersion("unknown"), "unknown"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.version.String(); got != c.expected {
				t.Errorf("String() = %q, want %q", got, c.expected)
			}
		})
	}
}

func TestAPIVersion_Validate(t *testing.T) {
	cases := []struct {
		name    string
		version APIVersion
		wantErr bool
	}{
		{"valid v1", APIVersionV1, false},
		{"invalid version", APIVersion("invalid"), true},
		{"empty version", APIVersion(""), true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.version.Validate()
			if (err != nil) != c.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, c.wantErr)
			}
			if c.wantErr && err != nil {
				exp := fmt.Sprintf("invalid API version: %s", c.version)
				if err.Error() != exp {
					t.Errorf("Validate() error = %q, want %q", err.Error(), exp)
				}
			}
		})
	}
}

func TestAPIVersion_GetBasePath(t *testing.T) {
	cases := []struct {
		name     string
		version  APIVersion
		expected string
	}{
		{"v1", APIVersionV1, "/api/v1"},
		{"unknown", APIVersion("foo"), "/api/foo"},
		{"empty", APIVersion(""), "/api/"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.version.GetBasePath(); got != c.expected {
				t.Errorf("GetBasePath() = %q, want %q", got, c.expected)
			}
		})
	}
}
