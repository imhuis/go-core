package main

import (
	"go-core/complex-type/htmltemplate"
	"go-core/complex-type/json-type"
	"os"
	"text/template"
)

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": htmltemplate.DaysAgo}).
	Parse(htmltemplate.Templ))

// repo:golang/go is:open json decoder
func main() {
	result, err := json_type.SearchIssues(os.Args[1:])
	if err = report.Execute(os.Stdout, result); err != nil {

	}
}
