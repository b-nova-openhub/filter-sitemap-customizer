package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchesPatternWithSingleWildcardWhitelistPattern(t *testing.T) {
	singleWhitelist := []string{"/home/content/*"}
	got := matchesPattern(singleWhitelist, "/home/content/blog.html", "")

	assert.Equal(t, true, got, "They should be equal!")
}

func TestMatchesPatternWithSingleWildcardBlacklistPattern(t *testing.T) {
	singleWhitelist := []string{"!/home/content/*"}
	got := matchesPattern(singleWhitelist, "/home/content/blog.html", "")

	assert.Equal(t, false, got, "They should be equal!")
}

func TestMatchesPatternWithMultipleWildcardWhitelistPatterns(t *testing.T) {
	whitelistPatterns := []string{"/home/content/*", "/home/service/*"}
	got := matchesPattern(whitelistPatterns, "/home/content/blog.html", "")

	assert.Equal(t, true, got, "They should be equal!")
}

func TestMatchesPatternWithMultipleDivergingPatterns(t *testing.T) {
	divergingPatterns := []string{"/home/content/*", "!/home/service"}
	got := matchesPattern(divergingPatterns, "/home/content/blog.html", "")

	assert.Equal(t, true, got, "They should be equal!")
}

func TestMatchesPatternWithMultipleConvergingPatterns(t *testing.T) {
	convergingPatterns := []string{"/home/content/*", "!/home/content/blog.html"}
	got := matchesPattern(convergingPatterns, "/home/content/blog.html", "")

	assert.Equal(t, true, got, "They should be equal!")
}
