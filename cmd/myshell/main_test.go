package main

import (
	"reflect"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

func TestTrimNewLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Remove Linux New Line (n)",
			input:    "Some text\n",
			expected: "Some text",
		},
		{
			name:     "Remove Windows New Line (rn)",
			input:    "Some text\r\n",
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

func TestGetParameters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Test Single Quotes",
			input:    "echo 'hello   world'",
			expected: []string{"echo", "hello   world"},
		},
		{
			name:     "Test Double Quotes",
			input:    "echo \"Hello   World\"",
			expected: []string{"echo", "Hello   World"},
		},
		{
			name:     "Test Double Quotes in Single Quotes",
			input:    "echo 'hello \" world'",
			expected: []string{"echo", "hello \" world"},
		},
		{
			name:     "Test Single Quotes in Double Quotes",
			input:    "echo \"hello ' world\"",
			expected: []string{"echo", "hello ' world"},
		},
		{
			name:     "Trim Space",
			input:    "world  example testscript",
			expected: []string{"world", "example", "testscript"},
		},
		{
			name:     "Test Multiple escaped words",
			input:    "\"world  example\"  \"test\"\"script",
			expected: []string{"world  example", "testscript"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := utils.ParseArguments(tt.input)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
