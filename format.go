package main

import "time"
import "strings"
import "fmt"

func formatMonth(m time.Month) string {
	s := m.String()

	s = strings.Replace(s, "January", "Enero", -1)
	s = strings.Replace(s, "February", "Febrero", -1)
	s = strings.Replace(s, "March", "Marzo", -1)
	s = strings.Replace(s, "April", "Abril", -1)
	s = strings.Replace(s, "May", "Mayo", -1)
	s = strings.Replace(s, "June", "Junio", -1)
	s = strings.Replace(s, "July", "Julio", -1)
	s = strings.Replace(s, "August", "Agosto", -1)
	s = strings.Replace(s, "September", "Septiembre", -1)
	s = strings.Replace(s, "October", "Octubre", -1)
	s = strings.Replace(s, "November", "Noviembre", -1)
	s = strings.Replace(s, "December", "Diciembre", -1)

	return s
}

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
