package filter

import (
	"github.com/spf13/viper"
	"github.com/yterajima/go-sitemap"
	"log"
	"path"
	"strings"
)

func Filter(url sitemap.URL, basePath string) string {
	filters := splitByComma(viper.GetString("filter"))
	if matchesPattern(filters, url.Loc, basePath) {
		return url.Loc
	}
	return ""
}

func splitByComma(s string) []string {
	return strings.Split(s, ",")
}

func matchesPattern(patterns []string, s, basePath string) bool {
	for _, p := range patterns {
		if !strings.HasPrefix(p, "!") {
			match, err := path.Match(basePath+p, s)
			if err != nil {
				log.Fatal(err)
			}
			if match == true {
				return match
			}
		} else {
			match, err := path.Match(basePath+strings.TrimPrefix(p, "!"), s)
			if err != nil {
				log.Fatal(err)
			}
			if match == false {
				return match
			}
		}
	}
	return false
}
