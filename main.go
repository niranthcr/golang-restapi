package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niranthcr/simple-web-api/models"
)

var Articles []models.Article

func getArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: /articles")
	json.NewEncoder(w).Encode(Articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: /article")
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range Articles {
		if article.ID == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var article models.Article
	json.Unmarshal(body, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var a models.Article
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &a)
	a.ID = id
	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			Articles = append(Articles, a)
		}
	}
	json.NewEncoder(w).Encode(a)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/articles", getArticles).Methods(http.MethodGet)
	router.HandleFunc("/article/{id}", getArticle).Methods(http.MethodGet)
	router.HandleFunc("/article", createArticle).Methods(http.MethodPost)
	router.HandleFunc("/article/{id}", deleteArticle).Methods(http.MethodDelete)
	router.HandleFunc("/article/{id}", updateArticle).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8000", router))

	// http.HandleFunc("/articles", getArticles)
	// log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {

	Articles = []models.Article{
		models.Article{ID: "1", Title: "Hello", Description: "Article Description", Content: "Article Content"},
		models.Article{ID: "2", Title: "Hello 2", Description: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Rest API v2.0 - Mux Routers")

	handleRequests()
}
