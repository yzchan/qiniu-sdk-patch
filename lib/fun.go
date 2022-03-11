package lib

import "time"

const layout = "20060102150405"

func formatDate(date string, add int) string {
	tmp, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}
	return tmp.AddDate(0, 0, add).Format(layout)
}

func FromDate(date string) string {
	return formatDate(date, 0)
}

func ToDate(date string) string {
	return formatDate(date, 1)
}
