package farmerly

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	parseView(w, "farmers.gohtml", r)
}

func Category(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := databaseConn()
		_, err := db.Query("SELECT * FROM categories")
		if err != nil {
			panic(err.Error())
		}
		parseView(w, "category.gohtml", r)
	} else {
		_ = r.ParseForm()
		if hasEmptyValues(r.Form) {
			log.Println("Is empty!")
		}
		http.Redirect(w, r, "/show", 301)
	}

}

func ViewGen(w http.ResponseWriter, r *http.Request) {
	parseView(w, "generated_content.gohtml", r)
}
