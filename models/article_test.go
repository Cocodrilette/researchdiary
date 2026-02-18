package models

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNewArticleFromTerminal(t *testing.T) {
	t.Run("all parameters", func(t *testing.T) {
		articleManager := ArticleManager{}

		input := "Test1\nAuthorLN\nAuthorFN\n2023-01-01\n1-10\nhttp://example.com\nJournal Name\nSome annotation\n"
		r := strings.NewReader(input)
		got, err := articleManager.NewFromTerminal(r)
		if err != nil {
			t.Fatalf("NewFromTerminal failed: %v", err)
		}

		url := "http://example.com"
		journal := "Journal Name"
		annot := "Some annotation"

		want := Article{
			Title:         "Test1",
			Author:        Author{FirstName: "AuthorFN", LastName: "AuthorLN"},
			DatePublished: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			PageRange:     &PageRange{1, 10},
			URL:           &url,
			JournalName:   &journal,
			Anotation:     &annot,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("invalid input", func(t *testing.T) {
		articleManager := ArticleManager{}

		input := "\nAuthorLN\nAuthorFN\n2023-01-01\n1-10\nhttp://example.com\nJournal Name\nSome annotation\n"
		r := strings.NewReader(input)
		_, err := articleManager.NewFromTerminal(r)
		if err == nil {
			t.Error("expected error for empty title")
		}

		input = "Test1\nAuthorLN\nAuthorFN\ninvalid-date\n1-10\nhttp://example.com\nJournal Name\nSome annotation\n"
		r = strings.NewReader(input)
		_, err = articleManager.NewFromTerminal(r)
		if err == nil {
			t.Error("expected error for invalid date")
		}

		input = "Test1\nAuthorLN\nAuthorFN\n2023-01-01\n1-10\ninvalid-url\nJournal Name\nSome annotation\n"
		r = strings.NewReader(input)
		_, err = articleManager.NewFromTerminal(r)
		if err == nil {
			t.Error("expected error for invalid URL")
		}
	})
}
