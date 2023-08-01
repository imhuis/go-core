package htmltemplate

import (
	"fmt"
	"text/template"
	"time"
)

const Templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}
`

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func Parse() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": DaysAgo}).
		Parse(Templ)
	if err != nil {
		panic(err)
	}
	fmt.Print(report)
}
