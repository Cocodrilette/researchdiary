package models

import (
	"testing"
	"time"
)

func TestAPAFormater(t *testing.T) {
	author := Author{FirstName: "Juan", LastName: "Moliner"}
	title := "Algunos problemas éticos de las tecnologías militares emergentes"
	pageRange := PageRange{522, 541}
	datePublished := time.Date(2018, 2, 19, 0, 0, 0, 0, time.UTC)
	url := "https://dialnet.unirioja.es/descarga/articulo/6467952.pdf"
	journalName := "Instituto Español de Estudio Estrátegicos"

	t.Run("APA with all data", func(t *testing.T) {
		article := Article{
			Author:        author,
			Title:         title,
			PageRange:     &pageRange,
			DatePublished: datePublished,
			URL:           &url,
			JournalName:   &journalName,
		}

		got := article.APA()
		want := "Moliner, J. (2018). Algunos problemas éticos de las tecnologías militares emergentes. Instituto Español de Estudio Estrátegicos, 522-541. https://dialnet.unirioja.es/descarga/articulo/6467952.pdf"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("APA with missing data", func(t *testing.T) {

		article := Article{
			Author:        author,
			Title:         title,
			PageRange:     &pageRange,
			DatePublished: datePublished,
		}
		got := article.APA()
		want := "Moliner, J. (2018). Algunos problemas éticos de las tecnologías militares emergentes. 522-541."

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
