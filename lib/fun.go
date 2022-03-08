package lib

import "time"

// TransDate 将"YYYY-MM-DD"格式的时间转化成 "20060102150405"格式
func TransDate(date string) string {
	tmp, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}
	return tmp.Format("20060102150405")
}
