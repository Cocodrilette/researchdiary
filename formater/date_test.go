package formater

import (
	"testing"
	"time"
)

func TestDateFormatter(t *testing.T) {
	publishDate := "2024-01-15"
	want := time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC)

	got, err := ParseString(publishDate)
	if err != nil {
		t.Errorf("%v", err)
	}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
