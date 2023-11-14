package main

import "time"
import "strings"
import "fmt"

var originalMonthNames = []string{
	"January", "February", "March", "April", "May", "June",
	"July", "August", "September", "October", "November", "December",
}

var originalDayNames = []string{
	"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
}

func formatMonth(m time.Month) string {
	s := m.String()

	for i := range originalMonthNames {
		s = strings.Replace(s, originalMonthNames[i], monthNames[i], -1)
	}

	return s
}

func formatTime(t time.Time, fmt string) string {
	s := t.Format(fmt)

	for i := range originalDayNames {
		s = strings.Replace(s, originalDayNames[i], dayNames[i], -1)
	}

	return s
}

func formatDates(
	t time.Time,
	all_day bool,
	end_date time.Time, end_time time.Time) string {
	var s string
	if all_day {
		s = formatTime(t, "Monday, 2/1/2006")
	} else {
		s = formatTime(t, "Monday, 2/1/2006 15:04")
	}

	y, m, d := time.Now().Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())

	y, m, d = t.Date()
	then := time.Date(y, m, d, 0, 0, 0, 0, t.Location())

	diff := int(then.Sub(today).Hours() / 24)

	if then.Before(today) {
		diff--
	}

	if diff > 0 {
		s += fmt.Sprintf(localeFormats["daysLeft"], diff)
	} else if diff < 0 {
		s += fmt.Sprintf(localeFormats["daysPast"], -diff)
	} else {
		s += localeFormats["today"]
	}

	if !end_date.IsZero() && !end_time.IsZero() {
		s += localeFormats["untilDate"]
		s += end_date.Format("2/1/2006")
		s += localeFormats["at"]
		s += end_time.Format("15:04")
	} else if !end_date.IsZero() {
		s += localeFormats["untilDate"] + end_date.Format("2/1/2006")
	} else if !end_time.IsZero() {
		s += localeFormats["untilTime"] + end_time.Format("15:04")
	}

	return s
}
