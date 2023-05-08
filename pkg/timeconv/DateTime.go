package timeconv

import (
	"time"
)

func ToDateOnly(datetime time.Time) string {
	return datetime.Format("02-01-2006")
}

func GetWeek(date time.Time) string {
	date = date.AddDate(0, 0, 7)
	return ToDateOnly(date)
}

func GetMonth(date time.Time) string {
	date = date.AddDate(0, 1, 0)
	return ToDateOnly(date)
}
