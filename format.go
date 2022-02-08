package main

import "time"
import "strings"
import "fmt"

func formatTime(t time.Time, fmt string) string {
	s := t.Format(fmt)

	s = strings.Replace(s, "Monday", "Lunes", -1)
	s = strings.Replace(s, "Tuesday", "Martes", -1)
	s = strings.Replace(s, "Wednesday", "Miércoles", -1)
	s = strings.Replace(s, "Thursday", "Jueves", -1)
	s = strings.Replace(s, "Friday", "Viernes", -1)
	s = strings.Replace(s, "Saturday", "Sábado", -1)
	s = strings.Replace(s, "Sunday", "Domingo", -1)

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
		s += fmt.Sprintf(" (quedan %d días)", diff)
	} else if diff < 0 {
		s += fmt.Sprintf(" (hace %d días)", -diff)
	} else {
		s += " (hoy)"
	}

	if !end_date.IsZero() && !end_time.IsZero() {
		s += " — hasta el día "
		s += end_date.Format("2/1/2006")
		s += " a las "
		s += end_time.Format("15:04")
	} else if !end_date.IsZero() {
		s += " — hasta el día " + end_date.Format("2/1/2006")
	} else if !end_time.IsZero() {
		s += " — hasta las " + end_time.Format("15:04")
	}

	return s
}
