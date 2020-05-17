package farmerly

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	parseEmptyView(w, "farmers.gohtml", r)
}


func Category(w http.Request, r *http.Request) {
	categories := fetchCategories()
	fmt.Println(categories)

}