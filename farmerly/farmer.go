package farmerly

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	parseEmptyView(w, "farmers.gohtml", r)
}


func Category(w http.ResponseWriter, r *http.Request) {
	categories := fetchCategories()
	parseView(w, "category.gohtml", categories, r)
}