package models

import "time"

type PageRange [2]int

type Article struct {
	Author        Author
	Title         string
	DatePublished time.Time
	PageRange     PageRange
	URL           *string
	JournalName   *string
	Anotation     *string
}
