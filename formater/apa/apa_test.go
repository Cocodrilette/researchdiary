package apa

import (
	"testing"
	"time"
)

func TestAPAFormater(t *testing.T) {
	url := "https://dialnet.unirioja.es/descarga/articulo/6467952.pdf"
	journalName := "Instituto Español de Estudio Estrátegicos"

	testArticle := Article{
		Author:        Author{firstName: "Juan", lastName: "Moliner"},
		Title:         "Algunos problemas éticos de las tecnologías militares emergentes",
		PageRange:     [2]int{522, 541},
		DatePublished: time.Date(2018, 2, 19, 0, 0, 0, 0, time.UTC),
		DateViewed:    time.Date(2026, 2, 18, 0, 0, 0, 0, time.UTC),
		URL:           &url,
		JournalName:   &journalName,
	}

	t.Run("APA with all data", func(t *testing.T) {
		article := testArticle

		got := article.APA()
		want := "Moliner, J. (2018). Algunos problemas éticos de las tecnologías militares emergentes. Instituto Español de Estudio Estrátegicos, 522-541. https://dialnet.unirioja.es/descarga/articulo/6467952.pdf"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("APA with missing data", func(t *testing.T) {
		article := Article{
			Author:        Author{firstName: "Juan", lastName: "Moliner"},
			Title:         "Algunos problemas éticos de las tecnologías militares emergentes",
			PageRange:     [2]int{522, 541},
			DatePublished: time.Date(2018, 2, 19, 0, 0, 0, 0, time.UTC),
			DateViewed:    time.Date(2026, 2, 18, 0, 0, 0, 0, time.UTC),
		}
		got := article.APA()
		want := "Moliner, J. (2018). Algunos problemas éticos de las tecnologías militares emergentes. 522-541."

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
