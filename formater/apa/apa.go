package apa

import (
	"strings"
	"text/template"
	"time"
)

type Author struct {
	firstName string
	lastName  string
}

func (a Author) LastName() string {
	return a.lastName
}

func (a Author) FirstInitial() string {
	if len(a.firstName) == 0 {
		return ""
	}
	return string([]rune(a.firstName)[0])
}

type PageRange [2]int

type Article struct {
	Author        Author
	Title         string
	DatePublished time.Time
	PageRange     PageRange
	URL           *string
	DateViewed    time.Time
	JournalName   string
	Anotation     string
}

const APAFormatTemplate = `{{.Author.LastName}}, {{.Author.FirstInitial}}. ({{.DatePublished.Year}}). {{.Title}}. {{.JournalName}}, {{index .PageRange 0}}-{{index .PageRange 1}}.{{if .URL}} {{.URL}}{{end}}`

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
