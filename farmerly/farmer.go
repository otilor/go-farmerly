package farmerly

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	parseView(w, "farmers.gohtml", r)
}
