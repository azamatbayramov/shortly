package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/azamatbayramov/shortly/pkg/link/validator"
)

func TestNetUrlValidator_Validate(t *testing.T) {
	v := validator.NetUrlValidator{}

	tests := []struct {
		link     string
		expected bool
	}{
		{"https://example.com", true},
		{"http://example.com", true},
		{"ftp://example.com", false},
		{"://example.com", false},
		{"example.com", false},
		{"", false},
		{"https://", false},
		{"https://sub.example.com", true},
		{"https://example.com/path", true},
		{"https://example.com?query=1", true},
		{"https://example.com#fragment", true},
		{"https://user:pass@example.com", true},
		{"http://localhost", true},
		{"http://127.0.0.1", true},
		{"https://.example.com", true},
		{"https://example.com:8080", true},
		{"https://-5-%-q.ca.x$1", false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, v.Validate(test.link))
	}
}
