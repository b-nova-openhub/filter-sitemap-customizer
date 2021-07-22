package filter

import (
	"github.com/yterajima/go-sitemap"
	"strings"
)

func Filter(url sitemap.URL, basePath string, filters []string) string {
	if matchesPattern(filters, url.Loc, basePath) {
		return url.Loc
	}
	return ""
}

func matchesPattern(patterns []string, s, basePath string) bool {
	for _, p := range patterns {
		return handleMatch(basePath, p, s)
	}
	return false
}

func handleMatch(basePath, pattern, s string) bool {
	if !strings.HasPrefix(pattern, "!") {
		whitelist := MatchApplier{WhitelistMatcher{}}
		if whitelist.Apply(s, basePath, pattern) {
			return true
		}
	} else {
		blacklist := MatchApplier{BlacklistMatcher{}}
		if blacklist.Apply(s, basePath, pattern) {
			return false
		}
	}
	return false
}
