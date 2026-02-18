package models

import (
	"strings"
	"text/template"
)

const APAFormatTemplate = `{{.Author.LastName}}, {{.Author.FirstInitial}}. ({{.DatePublished.Year}}). {{.Title}}.{{if .JournalName}} {{.JournalName}},{{end}} {{.PageRangeStart}}-{{.PageRangeEnd}}.{{if .URL}} {{.URL}}{{end}}`

func (a *Article) APA() string {
	tmpl, err := template.New("apa").Parse(APAFormatTemplate)
	if err != nil {
		return ""
	}

	var result strings.Builder
	err = tmpl.Execute(&result, a)
	if err != nil {
		return ""
	}

	return result.String()
}
