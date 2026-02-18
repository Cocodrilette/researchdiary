package apa

import (
	"fmt"
	"strings"
	"time"
)

type Author struct {
	firstName string
	lastName  string
}

type PageRange [2]int

type Article struct {
	Author        Author
	Title         string
	DatePublished time.Time
	PageRange     PageRange
	URL           string
	DateViewed    time.Time
	JournalName   string
	Anotation     string
}

const APAFormat = "{Author.lastName}, {Author.firstName.First()}. ({Self.DatePublished.Year}). {Self.Title}. {Self.JournalName}, {Self.PageRange[0]-Self.PageRange[1]}. {URL}"

func (a *Article) APA() string {

	APAStr := APAFormat

	APAStr = strings.Replace(APAStr, "{Author.lastName}", a.Author.lastName, 1)
	authorFirstNameFirstChar := string([]rune(a.Author.firstName)[0])
	APAStr = strings.Replace(APAStr, "{Author.firstName.First()}", authorFirstNameFirstChar, 1)
	APAStr = strings.Replace(APAStr, "{Self.DatePublished.Year}", fmt.Sprintf("%d", a.DatePublished.Year()), 1)
	APAStr = strings.Replace(APAStr, "{Self.Title}", a.Title, 1)
	APAStr = strings.Replace(APAStr, "{Self.JournalName}", a.JournalName, 1)

	pageRange := fmt.Sprintf("%d-%d", a.PageRange[0], a.PageRange[1])
	APAStr = strings.Replace(APAStr, "{Self.PageRange[0]-Self.PageRange[1]}", pageRange, 1)

	APAStr = strings.Replace(APAStr, "{URL}", a.URL, 1)

	return APAStr
}
