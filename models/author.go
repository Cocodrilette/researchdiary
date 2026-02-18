package models

type Author struct {
	FirstName string
	LastName  string
}

func (a Author) FirstInitial() string {
	if len(a.FirstName) == 0 {
		return ""
	}
	return string([]rune(a.FirstName)[0])
}
