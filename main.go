package main

import (
	"fmt"
	"time"

	"github.com/cocodrilette/researchdiary/models"
)

func main() {
	url := "https://dialnet.unirioja.es/descarga/articulo/6467952.pdf"
	journalName := "Instituto Español de Estudio Estrátegicos"

	article := models.Article{
		Author:        models.Author{FirstName: "Juan", LastName: "Moliner"},
		Title:         "Algunos problemas éticos de las tecnologías militares emergentes",
		PageRange:     [2]int{522, 541},
		DatePublished: time.Date(2018, 2, 19, 0, 0, 0, 0, time.UTC),
		DateViewed:    time.Date(2026, 2, 18, 0, 0, 0, 0, time.UTC),
		URL:           &url,
		JournalName:   &journalName,
	}

	fmt.Println(article.APA())
}
