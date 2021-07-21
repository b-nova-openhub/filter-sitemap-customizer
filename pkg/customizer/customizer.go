package customizer

import (
	"fmt"
	"github.com/b-nova-openhub/fisicus/pkg/filter"
	"github.com/yterajima/go-sitemap"
	"io/ioutil"
	"log"
	"net/http"
)

func GetContentHtml(url string) []string {
	content := make([]string, 0, 0)
	contentUrls := getFilteredUrls(url)
	for _, url := range contentUrls {
		get, getErr := http.Get(url)
		if getErr != nil {
			fmt.Println(getErr)
		}
		html, getErr := ioutil.ReadAll(get.Body)
		closeErr := get.Body.Close()
		if closeErr != nil {
			return nil
		}
		content = append(content, string(html))
	}
	return content
}

func getFilteredUrls(url string) []string {
	result := make([]string, 0, 0)
	urls := parseSitemap(url).URL
	for _, url := range urls {
		toAppend := filter.Filter(url)
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

