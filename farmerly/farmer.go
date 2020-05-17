package farmerly

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	parseEmptyView(w, "farmers.gohtml", r)
}


func Category(w http.Request, r *http.Request) {
	fetchCategories()
}