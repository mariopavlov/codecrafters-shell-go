package main

import "testing"

func TestTrimNewLine(t* testing.T) {
	tests := []struct {
		name string
		input string
		expected string
	} {
		{
			name: "Remove Linux New Line (n)",
			input: "Some text\n",
			expected: "Some text",
		},
		{
			name: "Remove Windows New Line (rn)",
			input: "Some text\r\n",
			expected: "Some text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := TrimNewLine(tt.input)

			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}