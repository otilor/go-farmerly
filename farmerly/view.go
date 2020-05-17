// Package farmerly implements a simple farming-algorithm.
// It maps farmers to content based on their selection

package farmerly

import (
	"html/template"
	"net/http"
)

func parseEmptyView(w http.ResponseWriter, tmpl string, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	isError(err)

	err = t.Execute(w, nil)
	isError(err)
}

func parseView(w http.ResponseWriter, tmpl string, pageVars []Categories, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	isError(err)

	err = t.Execute(w, pageVars)
	isError(err)
}

func ViewGen(w http.ResponseWriter, r *http.Request) {
	parseEmptyView(w, "generated_content.gohtml", r)
}
