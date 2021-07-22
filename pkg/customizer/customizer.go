package customizer

import (
	"fmt"
	"github.com/b-nova-openhub/fisicus/pkg/filter"
	"github.com/yterajima/go-sitemap"
	"log"
	"net/url"
)

func GetFilteredUrls(url string, filters []string) []string {
	result := make([]string, 0, 0)
	urls := parseSitemap(url).URL
	basePath := getBasePath(url)
	for _, url := range urls {
		toAppend := filter.Filter(url, basePath, filters)
		if toAppend != "" {
			result = append(result, toAppend)
			log.Println("Added following URL: ", toAppend)
		}
	}
	return result
}

func parseSitemap(url string) sitemap.Sitemap {
	smap, err := sitemap.Get(url, nil)
	if err != nil {
		fmt.Println(err)
	}
	for _, URL := range smap.URL {
		log.Println("Found following URL in sitemap.xml: ", URL.Loc)
	}
	return smap
}

func getBasePath(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u.Scheme + "://" + u.Host
}
