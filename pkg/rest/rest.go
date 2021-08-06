package rest

import (
	"encoding/json"
	"github.com/b-nova-openhub/fisicus/pkg/customizer"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sitemap", getSitemap).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), router))
}

func getSitemap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query()
	customizedUrls := customizer.GetFilteredUrls(getSitemapUrl(q), getFilters(q))
	json.NewEncoder(w).Encode(customizedUrls)
}

func getSitemapUrl(q url.Values) string {
	sm := q.Get("sitemap")
	if sm != "" {
		return sm
	}
	return viper.GetString("sitemap")
}

func getFilters(q url.Values) []string {
	f := q.Get("filter")
	if f != "" {
		return splitByComma(f)
	}
	return splitByComma(viper.GetString("filter"))
}

func splitByComma(s string) []string {
	return strings.Split(s, ",")
}
