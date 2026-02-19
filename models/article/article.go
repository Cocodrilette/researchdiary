package models

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/cocodrilette/researchdiary/formater"
	"github.com/cocodrilette/researchdiary/models"
	"gorm.io/gorm"
)

type TableName string

const tableName = TableName("articles")

type PageRange [2]int

type Article struct {
	gorm.Model

	AuthorID       uint
	Author         models.Author `gorm:"foreignKey:AuthorID"`
	Title          string
	DatePublished  time.Time
	PageRangeStart uint
	PageRangeEnd   uint
	URL            *string
	JournalName    *string
	Annotation     *string
}

type ArticleManager struct {
	DB *gorm.DB
}

type EmptyStrErr string

type UpdateQuery map[string]any

// Any type with an Error() string method fulfils the error interface
func (e EmptyStrErr) Error() string {
	return string(e)
}

func (a *ArticleManager) NewFromTerminal(in io.Reader) (Article, error) {
	reader := bufio.NewReader(in)
	var buffer bytes.Buffer

	title := getUserInput(reader, &buffer, "Ingresa el titulo: ")
	if strings.TrimSpace(title) == "" {
		return Article{}, EmptyStrErr("title")
	}

	authorLastname := getUserInput(reader, &buffer, "Ingresa el apellido del autor: ")
	if strings.TrimSpace(authorLastname) == "" {
		return Article{}, EmptyStrErr("author__last_name")
	}

	authorFirstname := getUserInput(reader, &buffer, "Ingresa el nombre del autor: ")

	if strings.TrimSpace(authorFirstname) == "" {
		return Article{}, EmptyStrErr("author__first_name")
	}

	dateStr := getUserInput(reader, &buffer, "Ingresa la fecha de publicacion (YYYY-MM-DD): ")
	var datePublished time.Time
	if dateStr != "" {
		if d, err := formater.ParseString(dateStr); err == nil {
			datePublished = d
		}
	}

	pageRangeStr := getUserInput(reader, &buffer, "Ingresa el rango de paginas (inicio-fin): ")
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

	pageRangeStartInt := uint(*pageRangeStart)
	pageRangeEndInt := uint(*pageRangeEnd)

	if pageRangeStartInt > pageRangeEndInt {
		return Article{}, fmt.Errorf("page range start cannot be greater than end")
	}

	urlStr := getUserInput(reader, &buffer, "Ingresa la URL: ")
	var url *string
	if urlStr != "" {
		url = &urlStr
	}

	journalStr := getUserInput(reader, &buffer, "Ingresa el nombre de la revista: ")
	var journalName *string
	if journalStr != "" {
		journalName = &journalStr
	}

	annotStr := getUserInput(reader, &buffer, "Ingresa la anotacion: ")
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
		Author:         models.Author{LastName: authorLastname, FirstName: authorFirstname},
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

func (a *ArticleManager) Find(query *Article) ([]Article, error) {
	var articles []Article

	result := a.DB.Where(query).Find(&articles)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find articles: %w", result.Error)
	}

	fmt.Printf("Found %d articles", result.RowsAffected)
	return articles, nil
}

func (a *ArticleManager) Save(article *Article) *gorm.DB {
	return a.DB.Save(article)
}

func (a *ArticleManager) Delete(article *Article) *gorm.DB {
	return a.DB.Delete(article)
}

func getUserInput(in *bufio.Reader, out io.Writer, instruction string) string {
	fmt.Fprintln(out, instruction)
	texto, err := in.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(texto)
}
