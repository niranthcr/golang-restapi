package tags

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/niranthcr/simple-web-api/pkg/db"
)

type Tag struct {
	Id   int    `json:id`
	Name string `json:name`
}

func GetTags(w http.ResponseWriter, r *http.Request) {
	d, err := db.GetDb()
	if err != nil {
		log.Fatal(err)
		return
	}
	var database = db.DbService{d}
	data, err := database.GetTagData()
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(w).Encode(data)

}
