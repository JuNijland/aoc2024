package day03

import "testing"

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "Simple multiplication",
			input: []string{
				"mul(2,3)",
			},
			expected: 6,
		},
		{
			name: "Corrupted input with space",
			input: []string{
				"mul(2, 3)",
			},
			expected: 0,
		},
		{
			name: "Corrupted input with missing closing parenthesis",
			input: []string{
				"mul(2,3",
			},
			expected: 0,
		},
		{
			name: "Multiple operations on one line",
			input: []string{
				"mul(2,3)mul(4,5)",
			},
			expected: 26,
		},
		{
			name: "Multiple lines",
			input: []string{
				"mul(2,3)",
				"mul(4,5)",
			},
			expected: 26,
		},
		{
			name: "Operations with text between them",
			input: []string{
				"start mul(2,3) middle mul(4,5) end",
			},
			expected: 26,
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: 0,
		},
		{
			name: "No operations",
			input: []string{
				"no operations here",
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_regex", func(t *testing.T) {
			result := solvePart1(tt.input)
			if result != tt.expected {
				t.Errorf("solvePart1Regex() = %v, want %v", result, tt.expected)
			}
		})
	}
}
