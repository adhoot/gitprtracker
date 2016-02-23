package ui

import (
	"fmt"
	"net/http"
	"html/template"
)

var report = template.Must(template.New("pullrequests").Parse(`
	<h1>{{.TotalCount}} Pull Requests </h1>
`))

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Git PR Tracker Main Page")
}
