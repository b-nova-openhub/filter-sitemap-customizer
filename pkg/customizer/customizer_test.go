package customizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBasePathWithFullyQualifiedUrl(t *testing.T) {
	got := getBasePath("https://golang.org/doc/install")
	assert.Equal(t, "https://golang.org", got, "They should be equal!")
}

func TestGetBasePathWithWorldWideWebUrl(t *testing.T) {
	got := getBasePath("https://www.golang.org/doc/install")
	assert.Equal(t, "https://www.golang.org", got, "They should be equal!")
}
