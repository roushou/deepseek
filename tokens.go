package deepseek

import (
	"math"
	"unicode"
)

// estimateTokens estimates the tokens for a given string according to the documentation https://api-docs.deepseek.com/quick_start/token_usage
func estimateTokens(text string) int64 {
	var tokens float64

	for _, rune := range text {
		switch {
		// English character
		case unicode.Is(unicode.Latin, rune):
			tokens += 0.3
		// Chinese character
		case unicode.Is(unicode.Han, rune):
			tokens += 0.6
		case unicode.IsDigit(rune):
			// Assuming a digit is considered an english character
			tokens += 0.3
		case unicode.IsSpace(rune):
			// The documentation doesn't specify this case
			continue
		case unicode.IsPunct(rune):
			// The documentation doesn't specify this case
			continue
		// Skipping...
		default:
			// The documentation doesn't specify this case
			continue
		}
	}

	// Rounding up since actual tokenization might split characters into subwords
	return int64(math.Ceil(tokens))
}
