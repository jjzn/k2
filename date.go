package main

import "time"

type YearWeek struct {
	Year int
	Week int
}

func isInDateRange(i Item, t time.Time) bool {
	var year_end, day_end int
	var month_end time.Month

	var year, day int
	var month time.Month

	year = i.Date.Year()
	month = i.Date.Month()
	day = i.Date.Day()

	if !i.EndDate.IsZero() {
		year_end = i.EndDate.Year()
		month_end = i.EndDate.Month()
		day_end = i.EndDate.Day()
	} else {
		year_end = i.Date.Year()
		month_end = i.Date.Month()
		day_end = i.Date.Day()
	}

	start := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	end := time.Date(year_end, month_end, day_end, 0, 0, 0, 0, t.Location())

	after := t.After(start) || t.Equal(start)
	before := t.Before(end) || t.Equal(end)
	return after && before
}

func weekBefore(week YearWeek) YearWeek {
	d := firstDayOfWeek(week.Year, week.Week)
	y, w := d.AddDate(0, 0, -7).ISOWeek()
	return YearWeek{y, w}
}

func weekAfter(week YearWeek) YearWeek {
	d := firstDayOfWeek(week.Year, week.Week)
	y, w := d.AddDate(0, 0, 7).ISOWeek()
	return YearWeek{y, w}
}

func firstDayOfWeek(year, week int) time.Time {
	// calculate iteratively

	// July 1st of the specified year is guaranteed to be in a week of `year`
	d := time.Date(year, 6, 1, 0, 0, 0, 0, time.Now().Location())
	_, wk := d.ISOWeek()

	d = d.AddDate(0, 0, 7*(week-wk))

	day := (int(d.Weekday()) + 6) % 7 // normalise so monday = 0 (week starts on monday)
	d = d.AddDate(0, 0, 1-day)        // shift date back so it is a monday

	return d
}
