package formater

import "time"

// Format date string with time.DateOnly format to time.Time
func ParseString(dateString string) (time.Time, error) {

	t, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
