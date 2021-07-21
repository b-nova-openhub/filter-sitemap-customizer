package rest

import (
	"encoding/json"
	"github.com/b-nova-openhub/fisicus/pkg/customizer"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", getHello).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getSitemap())
}

func getSitemap() []string {
	return customizer.GetFilteredUrls(viper.GetString("sitemap"))
}
