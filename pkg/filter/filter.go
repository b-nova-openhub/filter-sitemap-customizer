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
	var valid bool
	for _, p := range patterns {
		if !strings.HasPrefix(p, "!") {
			whitelist := MatchApplier{WhitelistMatcher{}}
			valid = whitelist.Apply(s, basePath, p)
		} else {
			blacklist := MatchApplier{BlacklistMatcher{}}
			valid = blacklist.Apply(s, basePath, p)
		}
		if valid {
			break
		}
	}
	return valid
}
