package main

import (
	"learning-go/complex-type/htmltemplate"
	"learning-go/complex-type/json1"
	"os"
	"text/template"
)

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": htmltemplate.DaysAgo}).
	Parse(htmltemplate.Templ))

// repo:golang/go is:open json decoder
func main() {
	result, err := json1.SearchIssues(os.Args[1:])
	if err = report.Execute(os.Stdout, result); err != nil {

	}
}
