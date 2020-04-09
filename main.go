package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/niranthcr/simple-web-api/pkg/articles"
	"github.com/niranthcr/simple-web-api/pkg/tags"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/articles", articles.GetArticles).Methods(http.MethodGet)
	router.HandleFunc("/article/{id}", articles.GetArticle).Methods(http.MethodGet)
	router.HandleFunc("/article", articles.CreateArticle).Methods(http.MethodPost)
	router.HandleFunc("/article/{id}", articles.DeleteArticle).Methods(http.MethodDelete)
	router.HandleFunc("/article/{id}", articles.UpdateArticle).Methods(http.MethodPut)

	router.HandleFunc("/tags", tags.GetTags).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", router))

	// http.HandleFunc("/articles", getArticles)
	// log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	handleRequests()
}
