package models

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/cocodrilette/researchdiary/formater"
	"gorm.io/gorm"
)

type PageRange [2]int

type Article struct {
	gorm.Model

	ID uint64 `gorm:"primaryKey;autoIncrement:true"`

	AuthorID       uint
	Author         Author `gorm:"foreignKey:AuthorID"`
	Title          string
	DatePublished  time.Time
	PageRangeStart uint64
	PageRangeEnd   uint64
	URL            *string
	JournalName    *string
	Annotation     *string
}

type ArticleManager struct {
	DB *gorm.DB
}

// const (
// 	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
// 	ErrAlreadyExists = DictionaryErr("a value for this key already exists")
// 	ErrDoesNotExists = DictionaryErr("does not exists key")
// )

type EmptyStrErr string

// Any type with an Error() string method fulfils the error interface
func (e EmptyStrErr) Error() string {
	return string(e)
}

func (a *ArticleManager) NewFromTerminal(in io.Reader) (Article, error) {
	reader := bufio.NewReader(in)

	title := getUserInput(reader, "Ingresa el titulo: ")
	if strings.TrimSpace(title) == "" {
		return Article{}, EmptyStrErr("title")
	}

	authorLastname := getUserInput(reader, "Ingresa el apellido del autor: ")
	if strings.TrimSpace(authorLastname) == "" {
		return Article{}, EmptyStrErr("author__last_name")
	}

	authorFirstname := getUserInput(reader, "Ingresa el nombre del autor: ")

	if strings.TrimSpace(authorFirstname) == "" {
		return Article{}, EmptyStrErr("author__first_name")
	}

	dateStr := getUserInput(reader, "Ingresa la fecha de publicacion (YYYY-MM-DD): ")
	var datePublished time.Time
	if dateStr != "" {
		if d, err := formater.ParseString(dateStr); err == nil {
			datePublished = d
		}
	}

	pageRangeStr := getUserInput(reader, "Ingresa el rango de paginas (inicio-fin): ")
	var pageRangeStart *int
	var pageRangeEnd *int

	if pageRangeStr != "" {
		parts := strings.Split(pageRangeStr, "-")
		if len(parts) == 2 {
			start, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			end, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

			if err1 == nil && err2 == nil {
				pageRangeStart = &start
				pageRangeEnd = &end
			}
		}
	}

	if pageRangeStart == nil || pageRangeEnd == nil {
		return Article{}, fmt.Errorf("invalid page range format, expected start-end")
	}

	pageRangeStartInt := uint64(*pageRangeStart)
	pageRangeEndInt := uint64(*pageRangeEnd)

	if pageRangeStartInt > pageRangeEndInt {
		return Article{}, fmt.Errorf("page range start cannot be greater than end")
	}

	urlStr := getUserInput(reader, "Ingresa la URL: ")
	var url *string
	if urlStr != "" {
		url = &urlStr
	}

	journalStr := getUserInput(reader, "Ingresa el nombre de la revista: ")
	var journalName *string
	if journalStr != "" {
		journalName = &journalStr
	}

	annotStr := getUserInput(reader, "Ingresa la anotacion: ")
	var annotation *string
	if annotStr != "" {
		annotation = &annotStr
	}

	if dateStr != "" && datePublished.IsZero() {
		return Article{}, fmt.Errorf("invalid date format, expected YYYY-MM-DD")
	}
	if urlStr != "" && !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		return Article{}, fmt.Errorf("URL must start with http:// or https://")
	}

	return Article{
		Title:          title,
		Author:         Author{LastName: authorLastname, FirstName: authorFirstname},
		DatePublished:  datePublished,
		PageRangeStart: pageRangeStartInt,
		PageRangeEnd:   pageRangeEndInt,
		URL:            url,
		JournalName:    journalName,
		Annotation:     annotation,
	}, nil
}

func (a *ArticleManager) Create(article *Article) error {
	err := a.DB.Create(article).Error
	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}

	fmt.Println("Created article with ID:", article.ID)
	return nil
}

func (a *ArticleManager) Find(query string) ([]Article, error) {
	var articles []Article
	result := a.DB.Find(&articles).Scan(&articles)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find articles: %w", result.Error)
	}

	fmt.Printf("Found %d articles", result.RowsAffected)
	return articles, nil
}

func getUserInput(in *bufio.Reader, instruction string) string {
	fmt.Print(instruction)
	texto, err := in.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(texto)
}
