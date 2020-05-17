package farmerly

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	parseEmptyView(w, "farmers.gohtml", r)
}


func Category(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		categories := fetchCategories()
		parseView(w, "category.gohtml", categories, r)
	}else {
		r.ParseForm()
		details := r.Form
		addUser(details)

	}

}