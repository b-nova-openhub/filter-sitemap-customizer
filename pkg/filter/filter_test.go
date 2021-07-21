package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchesPattern(t *testing.T) {
	patterns := []string{"/home/content/*"}
	got := matchesPattern(patterns, "/home/content/blog.html", "")

	assert.Equal(t, got, true, "They should be equal!")
}
