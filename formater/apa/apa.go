package apa

import (
	"strings"
	"text/template"
	"time"
)

type Author struct {
	FirstName string
	LastName  string
}

func (a Author) FirstInitial() string {
	if len(a.FirstName) == 0 {
		return ""
	}
	return string([]rune(a.FirstName)[0])
}

type PageRange [2]int

type Article struct {
	Author        Author
	Title         string
	DatePublished time.Time
	PageRange     PageRange
	URL           *string
	DateViewed    time.Time
	JournalName   *string
	Anotation     *string
}

const APAFormatTemplate = `{{.Author.LastName}}, {{.Author.FirstInitial}}. ({{.DatePublished.Year}}). {{.Title}}.{{if .JournalName}} {{.JournalName}},{{end}} {{index .PageRange 0}}-{{index .PageRange 1}}.{{if .URL}} {{.URL}}{{end}}`

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
