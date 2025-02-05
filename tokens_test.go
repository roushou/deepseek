package deepseek

import (
	"fmt"
	"testing"
)

func TestEstimateTokens(t *testing.T) {
	// Test cases for different scenarios
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "English text",
			input:    "Hello World",
			expected: 3, // 0.3 * 10 characters, rounded up
		},
		{
			name:     "Chinese text",
			input:    "你好世界",
			expected: 3, // 0.6 * 4 characters, rounded up
		},
		{
			name:     "Mixed text",
			input:    "Hello 世界",
			expected: 3, // 0.3 * 5 + 0.6 * 2, rounded up
		},
		{
			name:     "Digits",
			input:    "12345",
			expected: 2, // 0.3 * 5, rounded up
		},
		{
			name:     "Spaces and punctuation",
			input:    "  !@#$%^&*() ",
			expected: 0, // No tokens for spaces and punctuation
		},
		{
			name:     "Mixed with digits and symbols",
			input:    "Hello 世界 123 !@#",
			expected: 4, // 0.3 * 5 + 0.6 * 2 + 0.3 * 3, rounded up
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := estimateTokens(tt.input)
			if result != tt.expected {
				t.Errorf("estimateTokens(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}

	input := "你好世界"
	fmt.Println("wow")
	fmt.Println(estimateTokens(input))
}
