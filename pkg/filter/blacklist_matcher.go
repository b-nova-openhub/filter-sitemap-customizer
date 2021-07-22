package filter

import (
	"log"
	"path"
	"strings"
)

type BlacklistMatcher struct{}

func (b BlacklistMatcher) Match(s, prefix, pattern string) bool {
	m, err := path.Match(prefix+strings.TrimPrefix(pattern, "!"), s)
	if err != nil {
		log.Fatal(err)
	}
	return !m
}
