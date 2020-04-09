package articles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niranthcr/simple-web-api/models"
)

var Articles = []models.Article{
	models.Article{ID: "1", Title: "Hello", Description: "Article Description", Content: "Article Content"},
	models.Article{ID: "2", Title: "Hello 2", Description: "Article Description", Content: "Article Content"},
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: /articles")
	json.NewEncoder(w).Encode(Articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: /article")
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range Articles {
		if article.ID == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var article models.Article
	json.Unmarshal(body, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
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
