package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-JWT/auth"
)

type post struct {
	Title 	string `json:"title"`
	Tag		string `json:"tag"`
	URL		string `json:"url"`
}

func main() {
	r := mux.NewRouter()
	r.Handle("/public",public)
	r.Handle("/private", auth.JwtMiddleware.Handler(private))
	r.Handle("/auth", auth.GetTokenHandler)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "hoge",
        Tag:   "hoge",
        URL:   "https://hoge",
	}
	json.NewEncoder(w).Encode(post)
})

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: 	"piyo",
		Tag:	"piyo",
		URL:	"https://piyo",
	}
	json.NewEncoder(w).Encode(post)
})


