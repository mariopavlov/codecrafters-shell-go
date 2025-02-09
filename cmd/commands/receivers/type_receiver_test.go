package commands

import "testing"

func TestSearchCommandPath(t *testing.T) {
	tests := []struct {
		name     string
		found    bool
		input    string
		expected string
	}{
		{
			name:     "Look for zig on Path",
			found:    true,
			input:    "zig",
			expected: "/opt/homebrew/bin/zig",
		},
		{
			name:     "Look for windsurf on Path",
			found:    true,
			input:    "windsurf",
			expected: "/Users/mariopavlov/.codeium/windsurf/bin/windsurf",
		},
		{
			name:     "Look for cat on Path",
			found:    true,
			input:    "cat",
			expected: "/bin/cat",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, isFound := searchCommandPath(tt.input)

			if isFound != tt.found {
				t.Errorf("Expected found: %v, got: %v", tt.found, isFound)
			}

			if isFound && actual != tt.expected {
				t.Errorf("Expected path: %v, got: %v", tt.expected, actual)
			}
		})
	}
}
