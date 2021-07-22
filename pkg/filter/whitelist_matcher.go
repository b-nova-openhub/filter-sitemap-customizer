package filter

import (
	"log"
	"path"
)

type WhitelistMatcher struct{}

func (w WhitelistMatcher) Match(s, prefix, pattern string) bool {
	m, err := path.Match(prefix+pattern, s)
	if err != nil {
		log.Fatal(err)
	}
	return m
}
